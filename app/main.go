package main

import (
	"FifthProject/app/ws"
	"FifthProject/internal/utils"
)

func main() {

	utils.LogInit()
	ws.WsInit()

	defer utils.LogFile.Close()
}
