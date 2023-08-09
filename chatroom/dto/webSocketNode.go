package dto

import "github.com/gorilla/websocket"

type WebSocketNode struct {
	Connection *websocket.Conn
	DataQueue  chan []byte
}
