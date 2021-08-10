package ws

import (
	"FifthProject/internal/router"
	"FifthProject/internal/utils"
	"FifthProject/internal/ws"
	"net/http"
)

func WsInit() {

	/*router.WsRouter()

	err := http.ListenAndServe(":8899", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/

	//开一个goroutine执行开始程序
	go ws.Manager.Start()

	router.WsRouter()

	//注册默认路由为 /ws ，并使用wsHandler这个方法
	//http.HandleFunc("/ws", wsHandler)
	//监听本地的8011端口
	wsPort := utils.GetVal("server", "WsPort")
	http.ListenAndServe(":"+wsPort, nil)
}
