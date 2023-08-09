package service

import (
	"mentor_app/chatroom/dto"
	"sync"
)

var SocketConnectionMap map[string]*dto.WebSocketNode = make(map[string]*dto.WebSocketNode, 0)
var rwLock sync.RWMutex
var conversationMemberMap map[string][]string = make(map[string][]string, 0)
