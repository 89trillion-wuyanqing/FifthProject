package ctrl

import (
	"FifthProject/internal/service"
	"FifthProject/internal/utils"
	"FifthProject/internal/ws"
	"fmt"
	"log"
	"net/http"
)

func WsCtrl(res http.ResponseWriter, req *http.Request) {
	// 从请求里获取用户name
	username := req.Header.Get("username")
	fmt.Println(username + "用户连接")
	log.Println("INFO:" + username + "用户连接")
	//将http协议升级成websocket协议

	conn, err := utils.Upgraderdd.Upgrade(res, req, nil)
	if err != nil {

		http.NotFound(res, req)
		log.Println("ERROR:http错误，无法建立websocket长链接")
		return
	}
	//每一次连接都会新开一个client，client.id通过uuid生成保证每次都是不同的
	client := &ws.Client{Username: username, Socket: conn, Send: make(chan []byte)}
	service.NewUserRegister(client)
}
