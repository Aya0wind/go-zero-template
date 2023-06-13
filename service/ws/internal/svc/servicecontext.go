package svc

import (
	"book/service/ws/internal/config"
	"book/service/ws/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                config.Config
	WebSocketQueueChannel map[string]chan int
	DataModel             model.DataModel
	PolicyModel           model.PolicyModel
	TunnelModel           model.TunnelModel
	UserModel             model.UserModel
	DataWithTunnelModel   model.DataWithTunnelModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataBaseUrl)
	rawDB, err := conn.RawDB()
	err = rawDB.Ping()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:              c,
		DataModel:           model.NewDataModel(conn),
		PolicyModel:         model.NewPolicyModel(conn),
		TunnelModel:         model.NewTunnelModel(conn),
		UserModel:           model.NewUserModel(conn),
		DataWithTunnelModel: model.NewDataWithTunnelModel(conn),
	}
}
