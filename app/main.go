package main

import (
	"FifthProject/app/ws"
	"FifthProject/internal/utils"
)

func main() {

	ws.WsInit()
	//关闭日志文件
	defer utils.LogFile.Close()
}
