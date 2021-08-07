package ctrl

import (
	"FifthProject/internal/handler"
	"FifthProject/internal/utils"
	"fmt"
	"log"
	"net/http"
)

func WsCtrl(res http.ResponseWriter, req *http.Request) {
	//req.ParseForm()
	//uid:=req.Form["username"][0]      // 从请求里获取用户id
	username := req.Header.Get("username")
	fmt.Println(username + "用户连接")
	log.Println("INFO:" + username + "用户连接")
	//将http协议升级成websocket协议
	/*conn, err := (&websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true }}).Upgrade(res, req, nil)*/

	conn, err := utils.Upgraderdd.Upgrade(res, req, nil)
	if err != nil {

		http.NotFound(res, req)
		log.Println("ERROR:http错误，无法建立websocket长链接")
		return
	}

	//每一次连接都会新开一个client，client.id通过uuid生成保证每次都是不同的
	client := &handler.Client{Username: username, Socket: conn, Send: make(chan []byte)}
	//注册一个新的链接
	handler.Manager.Register <- client
	//启动协程收web端传过来的消息
	go client.Read()
	//启动协程把消息返回给web端
	go client.Write()
}
