package router

import (
	"FifthProject/internal/ctrl"
	"net/http"
)

func WsRouter() {
	http.HandleFunc("/ws", ctrl.WsCtrl)

}
