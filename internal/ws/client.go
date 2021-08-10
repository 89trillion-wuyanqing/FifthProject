package ws

import (
	"FifthProject/internal/model"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

//客户端 Client
type Client struct {

	//用户id
	Username string
	//连接的socket
	Socket *websocket.Conn
	//发送的消息
	Send chan []byte
}

//定义客户端结构体的read方法
func (c *Client) Read() {

	defer func() {

		fmt.Println("读方法中 ，我要关闭了")
		log.Println("INFO:" + c.Username + "连接关闭，进行注销")
		Manager.Unregister <- c
		//c.Socket.Close()
	}()

	for {

		//读取消息
		_, message, err := c.Socket.ReadMessage()
		//fmt.Println("读到消息")
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			//log.Println(err)
			fmt.Println("有错误，进行注销，关闭连接")
			log.Println("ERROR:读取用户为" + c.Username + "的消息时出错，注销该连接")
			Manager.Unregister <- c
			//c.Socket.Close()
			break
		}
		//如果没有错误信息就把信息放入broadcast
		//jsonMessage, _ := proto.Marshal(&model.GeneralReward{Msg: string(message),Type: 1})
		//var generalMsg = &model.GeneralReward{}
		var generalMsg = &model.GeneralReward{}
		e := proto.Unmarshal(message, generalMsg)
		if e != nil {
			log.Println("ERROR:用户" + c.Username + "在读送消息时，protobuf反序列化消息出错，关闭连接返回")
			return
		}
		log.Println("INFO:用户" + c.Username + "向服务端发送了一条消息，Msg:" + generalMsg.Msg + ",Type:" + strconv.Itoa(int(generalMsg.Type)))

		if generalMsg.Type == 2 {

			jsonSTr, _ := proto.Marshal(&model.GeneralReward{Msg: "Pong", Type: 2, Username: c.Username})
			//fmt.Println("收到ping消息："+string(jsonSTr)+"发送pong")
			c.Send <- jsonSTr
			continue
		}

		if generalMsg.Type == 3 {
			fmt.Println(generalMsg.Username + "关闭连接：")
			log.Println("INFO:用户" + c.Username + "发送退出消息，注销该连接")
			Manager.Unregister <- c
			//close(c.Send)
			//delete(Manager.Clients, c)
			continue
		}

		if generalMsg.Type == 5 {
			//fmt.Println("接收到list消息，并转发")

			labelStr := ""
			for k, v := range Manager.Clients {
				if v == true {
					//该用户在线
					labelStr += k.Username + "\n"

				}
			}
			reward := &model.GeneralReward{Msg: labelStr, Username: c.Username, Type: 5}
			jsonSTr, _ := proto.Marshal(reward)
			c.Send <- jsonSTr
			/*select {
			case c.Send <- jsonSTr:
			default:
				close(c.Send)
				delete(Manager.Clients, c)
			}*/
			continue
		}

		Manager.Broadcast <- message
	}
}

func (c *Client) Write() {

	defer func() {
		Manager.Unregister <- c
		fmt.Println("写方法中 ，我要关闭了")
		log.Println("INFO:用户为" + c.Username + "的连接关闭注销")
		//c.Socket.Close()
	}()

	for {

		/*select {
		case message,_:= <- c.Send:
			err :=c.Socket.WriteMessage(websocket.TextMessage, message)
			if (err != nil){
				log.Println(err.Error())
				Manager.Unregister <- c

				return
			}
		}*/
		select {

		//从send里读消息
		case message, ok := <-c.Send:
			//如果没有消息
			if !ok {

				/*c.Socket.WriteMessage(websocket.CloseMessage, []byte{
				})*/
				break
			}
			//有消息就写入，发送给web端
			log.Println("INFO:服务端向用户" + c.Username + "向服务端发送了一条消息")

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}

}
