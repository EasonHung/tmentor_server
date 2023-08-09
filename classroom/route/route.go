package route

import (
	"encoding/json"
	"mentor/classroom/constants/ws_cmd"
	"mentor/classroom/controller/class_controller"
	"mentor/classroom/controller/info_controller"
	"mentor/classroom/controller/ws_controller"
	"mentor/classroom/dto"
	"mentor/classroom/mentor_redis"
	"mentor/classroom/middleware"
	"mentor/classroom/middleware/log"
	"mentor/classroom/service/class_channel_service"
	"mentor/classroom/service/class_channel_var"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Route(route *gin.Engine) {
	classroomRoute := route.Group("/websocket")
	{
		classroomRoute.GET("/", classroomWs)
		classroomRoute.GET("/classroomList", ws_controller.GetClassroomMap)
		classroomRoute.GET("/redisChannel/list", ws_controller.GetRedisChannelList)
	}

	classroomInfoRoute := route.Group("/info")
	{
		classroomInfoRoute.GET("/token", middleware.VerifyToken, info_controller.GetClassroomToken)                // done
		classroomInfoRoute.GET("/classroomList", middleware.VerifyToken, info_controller.GetClassroomList)         // done
		classroomInfoRoute.POST("/enroll", middleware.VerifyToken, info_controller.EnrollClassroom)                // done
		classroomInfoRoute.POST("/init", info_controller.InitUser)                                                 // done
		classroomInfoRoute.POST("classSetting/add", middleware.VerifyToken, info_controller.AddClassSetting)       // done
		classroomInfoRoute.POST("classSetting/update", middleware.VerifyToken, info_controller.UpdateClassSetting) // done
		classroomInfoRoute.POST("classSetting/delete", middleware.VerifyToken, info_controller.DeleteClassSetting) // done
		classroomInfoRoute.GET("classSetting", middleware.VerifyToken, info_controller.GetClassSetting)            // done
		classroomInfoRoute.GET("/status", info_controller.GetClassroomStatus)                                      // done
		classroomInfoRoute.GET("/", middleware.VerifyToken, info_controller.GetUserClassroom)                      // done
		classroomInfoRoute.GET("/studentCount", info_controller.GetUserStudentCount)                               // done
		classroomInfoRoute.GET("/userClassroomId", middleware.VerifyToken, info_controller.GetUserClassroomId)     // done. this api can get others classroomId
		classroomInfoRoute.GET("/classRecord", middleware.VerifyToken, info_controller.GetClassRecord)
	}

	classroomInternalRoute := route.Group("/internal/info")
	{
		classroomInternalRoute.GET("/token", info_controller.GetClassroomToken) // done
	}

	classInfoRoute := route.Group("/class/info")
	{
		classInfoRoute.GET("/lastClass", middleware.VerifyToken, class_controller.GetLastUnfinishedClass)
		classInfoRoute.POST("/initClass", class_controller.InitClass)
		classInfoRoute.POST("/reimburseClass", class_controller.ReimburseClass)
		classInfoRoute.GET("/startTime", class_controller.GetStartTime)
		classInfoRoute.GET("/panic", class_controller.Panic)
	}
}

func classroomWs(c *gin.Context) {
	var (
		conn *websocket.Conn
		err  error
	)
	// upgrade to websocket connection
	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}

	var applicationType string
	applicationTypeHeader, _ := c.GetQuery("applicationType")
	if applicationTypeHeader == "web" {
		applicationType = class_channel_var.WEB
	} else {
		applicationType = class_channel_var.APP
	}

	userId, _ := c.GetQuery("userId")
	// add wesocket connection node
	node := &dto.WebSocketNode{
		Connection: conn,
		DataQueue:  make(chan []byte, 50),
		LocatedClassroomId: "",
	}
	reflect.ValueOf(node)
	defer func(userId string) {
		err := class_channel_service.NotifyClosureAndCleanMap(userId, applicationTypeHeader, node.LocatedClassroomId)
		if err != nil {
			log.Logger.Errorf("%+v\n", err)
		}
		log.Logger.Info("close connect " + userId)
		conn.Close()
	}(userId)

	// 訂閱
	sub := mentor_redis.Client.Subscribe(class_channel_var.WS_REDIS_PREFIX + applicationType + userId)
	defer func(userId string, sub *redis.PubSub) {
		conn.Close()
		sub.Unsubscribe(class_channel_var.WS_REDIS_PREFIX + applicationType + userId)
		sub.Close()
		log.Logger.Info("dispatch")
	}(userId, sub)

	go listenAndSendToClient(sub, node)
	recieveAndPublish(node)
	log.Logger.Info("disconnect")
}

func listenAndSendToClient(sub *redis.PubSub, node *dto.WebSocketNode) {
	for msg := range sub.Channel() {
		err := node.Connection.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		if err != nil {
			log.Logger.Error("[system] error Occured during sending single chat message")
			return
		}
	}
	log.Logger.Info("stop listening")
}

func recieveAndPublish(senderNode *dto.WebSocketNode) {
	for {
		_, data, err := senderNode.Connection.ReadMessage()
		if err != nil {
			log.Logger.Error(err.Error())
			log.Logger.Error("[system] error Occured during recieving single chat message", err)
			return
		}

		// json unmarshal
		message := dto.Message{}
		if err := json.Unmarshal(data, &message); err != nil {
			log.Logger.Errorf("[system] error Occured during unmarshal message %+v", err)
			return
		}

		switch message.Cmd {
		// todo: close room
		case ws_cmd.OpenRoom:
			// todo: verify user token
			ws_controller.OpenClassroom(message, senderNode)
		case ws_cmd.Ask:
			ws_controller.RequireAccess(message)
		case ws_cmd.Reject:
			ws_controller.RejectAccess(message)
		case ws_cmd.Accept:
			ws_controller.AcceptAndSendClassroomToken(message)
		case ws_cmd.JoinRoom:
			ws_controller.JoinClassRoom(message, senderNode)
		case ws_cmd.CloseRoom:
			ws_controller.CloseClassroom(message, senderNode)
		case ws_cmd.LeaveRoom:
			ws_controller.LeaveClassroom(message, senderNode)
		case ws_cmd.InstantMessage:
			ws_controller.SendInstantMsg(message)

		// start class procedure
		case ws_cmd.AskAcceptance:
			ws_controller.ClassRequest(message)
		case ws_cmd.AcceptClass:
			ws_controller.ClassAccept(message)
		case ws_cmd.ClockOn:
			ws_controller.ClassClockOn(message)
		case ws_cmd.FinishClass:
			ws_controller.FinishClass(message)

		// webRTC
		case ws_cmd.Offer:
			ws_controller.SendWebrtcOffer(message)
		case ws_cmd.Answer:
			ws_controller.SendWebrtcAnswer(message)
		case ws_cmd.Candidate:
			ws_controller.SendWebrtcCandidate(message)

		case ws_cmd.Ping:
			ws_controller.Heartbeat(message)
		default:
			log.Logger.Warnf("[system] wrong cmd: %s", message.Cmd)
		}
	}
}
