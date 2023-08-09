package controller

import (
	"encoding/json"
	"fmt"
	"mentor_app/chatroom/constants/chatroom_variables"
	"mentor_app/chatroom/dto"
	"mentor_app/chatroom/errcode"
	"mentor_app/chatroom/internal_error"
	"mentor_app/chatroom/mentor_redis"
	"mentor_app/chatroom/middleware/log"
	"mentor_app/chatroom/service"
	"mentor_app/chatroom/utils/time_utils"
	"mentor_app/chatroom/vo"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var rwLock sync.RWMutex

func Route(route *gin.Engine) {
	chatRoute := route.Group("/chatroom/websocket")
	{
		chatRoute.GET("/", chat)
		chatRoute.GET("/test/sub", testRedisSub)
		chatRoute.GET("/test/pub", testRedisPub)
		chatRoute.GET("/test/check", testRedisCheck)
	}

	chatInfoRoute := route.Group("/chatroom/info")
	{
		chatInfoRoute.POST("/user/init", initUser)
		chatInfoRoute.GET("/unreadMessage/one", getUnreadedMessagesOfSpecificConversation)
		chatInfoRoute.POST("/readCursor/update", updateReadCursor)
		chatInfoRoute.POST("/conversation/new", addConversation)
		chatInfoRoute.GET("/conversation/list", getConversationList)
		chatInfoRoute.GET("/conversation/sync/message", getConversationSyncMessage)
		chatInfoRoute.GET("/readCursor/another", getAnotherCursor)
		chatInfoRoute.GET("/readCursor/self", getSelfCursor)
	}
}

func testRedisCheck(c *gin.Context) {
	// 没有指定查询channel的匹配模式，则返回所有的channel
	// 匹配user_开头的channel
	chs, _ := mentor_redis.Client.PubSubChannels("*").Result()
	for _, ch := range chs {
		log.Logger.Info(ch)
	}
}

func testRedisSub(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	sub := mentor_redis.Client.Subscribe(userId)
	for msg := range sub.Channel() {
		// 打印收到的消息
		fmt.Println(msg.Channel)
		fmt.Println(msg.Payload)
	}
}

func testRedisPub(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	mentor_redis.Client.Publish(userId, "message")
}

func getConversationSyncMessage(c *gin.Context) {
	conversationId, _ := c.GetQuery("conversationId")
	lastMessageId, _ := c.GetQuery("lastMessageId")

	messageDao, err := service.GetSyncMessages(conversationId, lastMessageId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	result := vo.GetConversationMessageRes{}
	result.ChatMessageListConvertor(messageDao)
	c.JSON(200, vo.NewSuccessResponse(result))
}

func updateReadCursor(c *gin.Context) {
	var request vo.UpdateReadCursorReq
	c.BindJSON(&request)

	err := service.UpdateReadCursor(request.ConversationId, request.UserId, request.DeviceId, request.LastMessageId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, "ok")
}

func getConversationList(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	conversationList, err := service.GetUserConversationList(userId)
	if err != nil {
		log.Logger.Error(err)
		c.JSON(500, err)
		return
	}
	result := vo.GetConversationListRes{}
	result.ConversationConvertor(conversationList)

	for index, conversationInfoVo := range result.Data {
		lastMessage, lastMessageTime, err := service.GetLastConversationMessageAndTime(conversationInfoVo.ConversationId)
		if err != nil {
			log.Logger.Error(err)
			c.JSON(500, err)
			return
		}
		conversationUserInfo, err := service.GetConversationUserInfo(conversationInfoVo.Participants, userId)
		if err != nil {
			log.Logger.Error(err)
			c.JSON(500, err)
			return
		}

		unReadedCount, err := service.CountUnReadedMessage(conversationInfoVo.ConversationId, userId)

		lastMessageTime, err = time_utils.TimeInTaipei(lastMessageTime)
		if err != nil {
			log.Logger.Error(err)
			c.JSON(500, err)
			return
		}
		result.Data[index].LastMessage = lastMessage
		result.Data[index].UnReadedCount = int(unReadedCount)
		result.Data[index].UserInfo = conversationUserInfo
		if lastMessageTime.String() == "0001-01-01 08:06:00 +0806 LMT" {
			result.Data[index].LastMessageTime = ""
		} else {
			result.Data[index].LastMessageTime = lastMessageTime.String()
		}
	}

	c.JSON(200, result.Data)
}

func initUser(c *gin.Context) {
	var request vo.InitUserReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Error(err)
		c.JSON(500, err)
		return
	}

	err = service.InitUser(request.UserId)
	if err != nil {
		log.Logger.Error(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, "OK")
}

func getAnotherCursor(c *gin.Context) {
	userId, _ := c.GetQuery("userId")
	conversationId, _ := c.GetQuery("conversationId")

	err, cursor := service.GetAnotherCursor(conversationId, userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(cursor))
}

func getSelfCursor(c *gin.Context) {
	userId, _ := c.GetQuery("userId")
	conversationId, _ := c.GetQuery("conversationId")

	err, cursor := service.GetSelfCursor(conversationId, userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(cursor))
}

func addConversation(c *gin.Context) {
	c.Next()
	var request dto.Conversation // 應該要搬到vo
	c.BindJSON(&request)

	// verify conversation type
	if request.Type == 0 {
		if len(request.Participants) > 2 {
			c.JSON(500, "cannot add single conversation with more than 2 people")
			return
		}
	}

	conversationId, err := service.AddConversation(request.UserId, request.Participants, request.Type)
	if err != nil {
		if errors.As(err, &internal_error.DuplicateAddConversationError{}) {
			log.Logger.Errorf("%+v\n", err)
			res := vo.NewErrorResponse(errcode.DuplicateAddConversation)

			res.Data = vo.AddConversationRes{
				ConversationId: conversationId,
			}
			c.JSON(200, res)
			return
		} else {
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	res := vo.AddConversationRes{
		ConversationId: conversationId,
	}
	c.JSON(200, vo.NewSuccessResponse(res))
	return
}

func getUnreadedMessagesOfSpecificConversation(c *gin.Context) {
	conversationId, _ := c.GetQuery("conversationId")
	userId, _ := c.GetQuery("userId")

	err, unreadMessages := service.GetUnreadMessageByConversationId(conversationId, userId)
	if err != nil {
		log.Logger.Error("[system] error occured when find unread messages")
		c.JSON(500, err)
		return
	}

	c.JSON(200, unreadMessages)
}

func chat(c *gin.Context) {
	var (
		conn *websocket.Conn
		err  error
	)
	// upgrade to websocket connection
	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	userId, _ := c.GetQuery("userId")
	deviceId, _ := c.GetQuery("deviceId")

	// add wesocket connection node
	node := &dto.WebSocketNode{
		Connection: conn,
		DataQueue:  make(chan []byte, 50),
	}
	rwLock.Lock()
	service.SocketConnectionMap[userId+deviceId] = node
	rwLock.Unlock()

	sub := mentor_redis.Client.Subscribe(chatroom_variables.WS_REDIS_PREFIX + userId)
	defer func(userId string, sub *redis.PubSub) {
		delete(service.SocketConnectionMap, userId+deviceId)
		conn.Close()
		sub.Unsubscribe(chatroom_variables.WS_REDIS_PREFIX + userId)
		sub.Close()
		log.Logger.Info("dispatch")
	}(userId, sub)

	go listenAndSend(sub, node)
	recieveAndDispatch(node, deviceId)
	log.Logger.Info("disconnect")
}

func listenAndSend(sub *redis.PubSub, node *dto.WebSocketNode) {
	for msg := range sub.Channel() {
		err := node.Connection.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		if err != nil {
			log.Logger.Error("[system] error Occured during sending single chat message")
			return
		}
	}
	log.Logger.Info("stop listening")
}

func recieveAndDispatch(senderNode *dto.WebSocketNode, deviceId string) {
	for {
		_, data, err := senderNode.Connection.ReadMessage()
		if err != nil {
			log.Logger.Error("[system] error Occured during recieving single chat message", err)
			return
		}

		// json unmarshal
		message := dto.Message{}
		if err := json.Unmarshal(data, &message); err != nil {
			log.Logger.Error("[system] error Occured during unmarshal message", err)
			return
		}

		switch Cmd(message.Cmd) {
		case SingleChat:
			service.SingleChatMessageProgress(senderNode, &message, data)
		case SingleChatReaded:
			service.SingleChatReadedProgress(message, deviceId, data)
		case GroupChat:
			service.GroupChatMessageProgress(senderNode, &message, data)
		case GroupChatReaded:
			service.GroupChatReadedProgress(&message, data)
		case ClassroomInfo:
			service.ClassroomInfoMessageProgress(senderNode, &message, data)
		case Heartbeat:
			service.HeartbeatProgress(senderNode, &message, data)
		default:
			log.Logger.Warn("[system] wrong cmd", message.Cmd)
		}
	}
}
