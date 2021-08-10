package ws

import (
	"FifthProject/internal/router"
	"FifthProject/internal/utils"
	"FifthProject/internal/ws"
	"net/http"
)

func WsInit() {

	//开一个goroutine执行开始程序
	go ws.Manager.Start()

	router.WsRouter()

	//监听本地的8001端口
	wsPort := utils.GetVal("server", "WsPort")
	http.ListenAndServe(":"+wsPort, nil)
}
