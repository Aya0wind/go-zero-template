package handler

import (
	"book/service/ws/internal/logic"
	"book/service/ws/internal/svc"
	"context"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpx"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func wsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		wsLogic := logic.NewWsLogic(context.Background(), svcCtx)
		err = wsLogic.Ws(c)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
