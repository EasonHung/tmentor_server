package class_channel_var

import (
	"sync"
)

var RwLock sync.RWMutex
var UserLocateMap map[string]string = make(map[string]string, 0) // 如果user進入房間就會出現在 map 裡

const WS_REDIS_PREFIX = "classroom-ws-"

// application type
const (
	WEB = "web-"
	APP = "app-"
)

// const JOIN_ROOM_SCRIPT = `
// 	local memberList = redis.call("lrange", KEYS[1], 0, -1)

// 	if( #memberList == 0) then
// 		return "class not open"
// 	elseif( #memberList > 2 ) then
// 		return "already in class"
// 	else
// 		redis.call('lpush', KEYS[1], KEYS[2])
// 		return "join class success"
// 	end

// 	return -1
// `
