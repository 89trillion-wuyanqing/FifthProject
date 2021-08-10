package service

import (
	"FifthProject/internal/ws"
)

/**
有新用户登陆 发送登陆消息 给所以用户发送list消息更新所有用户列表
*/
func NewUserRegister(client *ws.Client) {
	//获取所有的用户信息封装广播
	labelStr := ""
	for k, v := range ws.Manager.Clients {
		if v == true {
			//该用户在线
			labelStr += k.Username + "\n"

		}
	}

	//reward := &model.GeneralReward{Msg: labelStr, Username:"system", Type: 5}
	//jsonSTr, _ := proto.Marshal(reward)

	//注册一个新的链接
	ws.Manager.Register <- client
	//fmt.Println("11111")
	//广播
	//ws.Manager.Broadcast <- jsonSTr
	//fmt.Println("22222")
	//启动协程收web端传过来的消息
	go client.Read()
	//启动协程把消息返回给web端
	go client.Write()
	//fmt.Println("33333")
}
