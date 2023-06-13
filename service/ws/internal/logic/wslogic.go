package logic

import (
	"book/service/ws/internal/svc"
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) WsLogic {
	return WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsLogic) Ws(c *websocket.Conn) error {
	//connectionChannel := make(chan string)
	//l.svcCtx.WebSocketQueueChannel[]
	//for {
	//	c.
	//}
	return nil
}

func (l *WsLogic) verifyAuth(c *websocket.Conn) (err error) {
	_, _, err = c.ReadMessage()
	if err != nil {
		l.Info(err)
	}
	return
}
