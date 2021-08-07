package ws

import (
	"FifthProject/internal/handler"
	"FifthProject/internal/router"
	"net/http"
)

func WsInit() {

	/*router.WsRouter()

	err := http.ListenAndServe(":8899", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/

	//开一个goroutine执行开始程序
	go handler.Manager.Start()

	router.WsRouter()

	//注册默认路由为 /ws ，并使用wsHandler这个方法
	//http.HandleFunc("/ws", wsHandler)
	//监听本地的8011端口
	http.ListenAndServe(":8011", nil)
}
