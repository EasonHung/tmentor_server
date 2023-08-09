package ws_controller

import (
	"mentor/classroom/constants/ws_cmd"
	"mentor/classroom/dto"
	"mentor/classroom/dto/req"
	"mentor/classroom/mentor_redis"
	"mentor/classroom/middleware/log"
	"mentor/classroom/service/class_channel_service"
	"mentor/classroom/service/class_channel_var"
	"mentor/classroom/service/class_process_service"
	"mentor/classroom/service/class_time_service"
	"mentor/classroom/service/classroom_access_service"
	"mentor/classroom/service/classroom_message_service"
	"mentor/classroom/service/message_service"
	"mentor/classroom/service/web_rtc_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetRedisChannelList(c *gin.Context) {
	// 没有指定查询channel的匹配模式，则返回所有的channel
	// 匹配user_开头的channel
	chs, _ := mentor_redis.Client.PubSubChannels(class_channel_var.WS_REDIS_PREFIX + "*").Result()
	res := make([]string, 0)
	for _, ch := range chs {
		res = append(res, ch)
	}

	c.JSON(200, res)
	return
}

func GetClassroomMap(c *gin.Context) {
	c.JSON(200, class_channel_service.GetClassroomList())
}

func OpenClassroom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) {
	err := classroom_access_service.OpenClassroom(senderMessage, senderWsNode)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func RequireAccess(senderMessage dto.Message) {
	err := classroom_access_service.RequireAccess(senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func RejectAccess(senderMessage dto.Message) {
	err := classroom_access_service.RejectAccess(senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func AcceptAndSendClassroomToken(senderMessage dto.Message) {
	err := classroom_access_service.AcceptAndSendClassroomToken(senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func JoinClassRoom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) {
	err := classroom_access_service.JoinClassRoom(senderMessage, senderWsNode)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func SendInstantMsg(senderMessage dto.Message) {
	err := classroom_message_service.SendInstantMessage(senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func SendWebrtcOffer(senderMessage dto.Message) {
	err := web_rtc_service.SendOffer(senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func SendWebrtcAnswer(senderMessage dto.Message) {
	err := web_rtc_service.SendAnswer(senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func SendWebrtcCandidate(senderMessage dto.Message) {
	err := web_rtc_service.SendCandidate(senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func CloseClassroom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) {
	err := classroom_access_service.CloseClassroom(senderMessage, senderWsNode)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func LeaveClassroom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) {
	err := classroom_access_service.LeaveClassroom(senderMessage, senderWsNode)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func ClassRequest(senderMessage dto.Message) {
	err := class_process_service.AskForAcceptance(senderMessage.Message.(string), senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func ClassAccept(senderMessage dto.Message) {
	err := class_process_service.AcceptClass(senderMessage.Message.(string), senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func ClassClockOn(senderMessage dto.Message) {
	err, clockOnReq := req.JsonStrToClockOnReq(senderMessage.Message.(string))
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}

	err = class_time_service.ClockOn(clockOnReq.ClassId, senderMessage.SenderId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func FinishClass(senderMessage dto.Message) {
	err := class_process_service.FinishClass(senderMessage.Message.(string), senderMessage.ClassroomId, senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}

func Heartbeat(senderMessage dto.Message) {
	senderMessage.Cmd = ws_cmd.Pong
	err := message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		senderMessage.ToErrorCmd(err.Error())
		message_service.SendMessage(senderMessage.ApplicationType, senderMessage.SenderId, senderMessage)
		return
	}
}
