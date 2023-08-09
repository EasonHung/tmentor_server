package service

import (
	"encoding/json"
	"mentor_app/chatroom/constants/chatroom_variables"
	"mentor_app/chatroom/dao/chat_message_dao"
	"mentor_app/chatroom/dao/conversation_dao"
	"mentor_app/chatroom/dto"
	"mentor_app/chatroom/internalAPI/user_system_api"
	"mentor_app/chatroom/mentor_redis"
	"mentor_app/chatroom/middleware/log"
	"mentor_app/chatroom/service/notification_service"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

func SingleChatMessageProgress(node *dto.WebSocketNode, messageDto *dto.Message, data []byte) error {
	err, messageId := saveSingleChatMessage(messageDto)
	if err != nil {
		return err
	}

	err, senderInfo := user_system_api.GetUserInfo(messageDto.SenderId)
	if err != nil {
		log.Logger.Errorf("error push notification to receiver %+v", err)
	}

	messageDto.MessageId = messageId
	messageDto.SenderAvatarUrl = senderInfo.Data.AvatorUrl
	messageDto.Time = bson.ObjectIdHex(messageId).Time()

	jsonByte, _ := json.Marshal(messageDto)
	_, err = mentor_redis.Client.Publish(chatroom_variables.WS_REDIS_PREFIX+messageDto.SenderId, string(jsonByte)).Result()
	if err != nil {
		log.Logger.Error("error publish chat message to sender", err)
		return err
	}
	log.Logger.Info(chatroom_variables.WS_REDIS_PREFIX + messageDto.RecieverId)
	resReceiver, err := mentor_redis.Client.Publish(chatroom_variables.WS_REDIS_PREFIX+messageDto.RecieverId, string(jsonByte)).Result()
	if err != nil {
		log.Logger.Error("error publish chat message to receiver", err)
		return err
	}
	if resReceiver == 0 {
		err = notification_service.PushChatNotification(messageDto.RecieverId, messageDto.Message, messageDto.ConversationId, messageDto.SenderId, messageDto.MessageId, messageDto.Time.String(), senderInfo.Data.AvatorUrl)
		if err != nil {
			log.Logger.Errorf("error push notification to receiver %+v", err)
		}
	}

	return nil
}

func SingleChatReadedProgress(messageDto dto.Message, deviceId string, data []byte) error {
	err := conversation_dao.UpsertConversationCursor(messageDto.ConversationId, messageDto.SenderId, deviceId, messageDto.MessageId)
	if err != nil {
		log.Logger.Error("error update cursor", err)
	}

	jsonByte, _ := json.Marshal(messageDto)
	_, err = mentor_redis.Client.Publish(chatroom_variables.WS_REDIS_PREFIX+messageDto.RecieverId, string(jsonByte)).Result()
	if err != nil {
		log.Logger.Error("error publish read message to receiver", err)
	}
	_, err = mentor_redis.Client.Publish(chatroom_variables.WS_REDIS_PREFIX+messageDto.SenderId, string(jsonByte)).Result()
	if err != nil {
		log.Logger.Error("error publish read message to sender", err)
	}
	return nil
}

func saveSingleChatMessage(dto *dto.Message) (error, string) {
	// save message
	messageId := bson.NewObjectId().Hex()
	dto.Time = bson.ObjectIdHex(messageId).Time()
	message := chat_message_dao.ChatMessage{
		MessageId:      messageId,
		SenderId:       dto.SenderId,
		Data:           dto,
		ReadedBy:       []chat_message_dao.ReadedBy{},
		ReadCount:      0,
		ConversationId: dto.ConversationId,
	}
	_, err := chat_message_dao.InsertOne(&message)
	if err != nil {
		log.Logger.Error("[DB] error occured when insert message", err)
		return err, ""
	}

	return err, messageId
}

func HeartbeatProgress(node *dto.WebSocketNode, messageDto *dto.Message, data []byte) error {
	// 回傳pong給sender
	messageDto.Message = "pong"

	data, _ = json.Marshal(messageDto)

	// node.DataQueue <- data
	err := node.Connection.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Logger.Error("[system] error Occured during sending single chat message")
		return err
	}
	return nil
}
