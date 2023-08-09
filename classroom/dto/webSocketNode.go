package dto

import "github.com/gorilla/websocket"

type WebSocketNode struct {
	Connection         *websocket.Conn
	LocatedClassroomId string
	DataQueue          chan []byte // 目前沒用到
}
