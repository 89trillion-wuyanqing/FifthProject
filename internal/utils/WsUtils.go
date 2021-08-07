package utils

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	// 服务器应用程序从HTTP请求处理程序调用Upgrader.Upgrade方法以获取* Conn;
	Upgraderdd = websocket.Upgrader{
		// 读取存储空间大小
		ReadBufferSize: 1024,
		// 写入存储空间大小
		WriteBufferSize: 1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
