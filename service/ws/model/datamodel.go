package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	dataFieldNames          = builder.RawFieldNames(&Data{})
	dataRows                = strings.Join(dataFieldNames, ",")
	dataRowsExpectAutoSet   = strings.Join(stringx.Remove(dataFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	dataRowsWithPlaceHolder = strings.Join(stringx.Remove(dataFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	DataModel interface {
		Insert(data *Data) (sql.Result, error)
		FindOne(id int64) (*Data, error)
		Update(data *Data) error
		Delete(id int64) error
		FindOneByOwnerId(ownerid int64) (*Data, error)
	}

	defaultDataModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Data struct {
		Id         int64          `db:"id"`          // 自增id
		OwnerId    string         `db:"owner_id"`    // 所有者唯一id;使用UUID
		DataPath   sql.NullString `db:"data_path"`   // 数据路径;后续可能换成url
		UseCount   sql.NullInt64  `db:"use_count"`   // 使用计数
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"` // 更新时间
	}
)

func (m *defaultDataModel) FindOneByOwnerId(ownerid int64) (*Data, error) {
	query := fmt.Sprintf("select %s from %s where `owner_id` = ? limit 1", dataRows, m.table)
	var resp Data
	err := m.conn.QueryRow(&resp, query, ownerid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func NewDataModel(conn sqlx.SqlConn) DataModel {
	return &defaultDataModel{
		conn:  conn,
		table: "`Data`",
	}
}

func (m *defaultDataModel) Insert(data *Data) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, dataRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.OwnerId, data.DataPath, data.UseCount)
	return ret, err
}

func (m *defaultDataModel) FindOne(id int64) (*Data, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dataRows, m.table)
	var resp Data
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDataModel) Update(data *Data) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dataRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.OwnerId, data.DataPath, data.UseCount, data.Id)
	return err
}

func (m *defaultDataModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
