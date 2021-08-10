package ws

import (
	"FifthProject/internal/model"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

//客户端管理
type ClientManager struct {

	//客户端 map 储存并管理所有的长连接client，在线的为true，不在的为false
	Clients map[*Client]bool
	//web端发送来的的message我们用broadcast来接收，并最后分发给所有的client
	Broadcast chan []byte
	//新创建的长连接client
	Register chan *Client
	//新注销的长连接client
	Unregister chan *Client
}

//创建客户端管理者
var Manager = ClientManager{

	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

func (manager *ClientManager) Start() {

	for {

		select {

		//如果有新的连接接入,就通过channel把连接传递给conn
		case conn := <-manager.Register:
			//把客户端的连接设置为true
			manager.Clients[conn] = true
			//把返回连接成功的消息protobuf格式化
			//jsonMessage, _ := proto.Marshal(&model.GeneralReward{Msg: "新用户"+conn.Username+"登陆",Type: 1})
			jsonMessage, _ := proto.Marshal(&model.GeneralReward{Msg: "user is " + conn.Username + " login", Type: 1, Username: conn.Username})

			fmt.Println("发送登陆消息" + string(jsonMessage))
			log.Println("INFO：" + conn.Username + "发送登陆消息")
			//调用客户端的send方法，发送消息
			manager.send(jsonMessage)
			//如果连接断开了
		case conn := <-manager.Unregister:
			//判断连接的状态，如果是true,就关闭send，删除连接client的值
			if _, ok := manager.Clients[conn]; ok {

				close(conn.Send)
				delete(manager.Clients, conn)
				//jsonMessage, _ := proto.Marshal(&model.GeneralReward{Msg: "用户"+conn.Username+"退出",Type: 3})
				jsonMessage, _ := proto.Marshal(&model.GeneralReward{Msg: "user is " + conn.Username + " exit", Type: 3, Username: conn.Username})
				fmt.Println("关闭消息")
				log.Println("INFO：" + conn.Username + "发送登出消息")
				manager.send(jsonMessage)
			}
			//广播
		case message := <-manager.Broadcast:
			fmt.Println("广播")
			log.Println("INFO：发送广播消息")
			manager.send(message)

		}
	}
}

//定义客户端管理的send方法
func (manager *ClientManager) send(message []byte) {

	for conn := range manager.Clients {

		conn.Send <- message

	}
}
