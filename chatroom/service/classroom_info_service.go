package service

import (
	"encoding/json"
	"mentor_app/chatroom/constants/chatroom_variables"
	"mentor_app/chatroom/dto"
	"mentor_app/chatroom/internalAPI/classroom_system_api"
	"mentor_app/chatroom/mentor_redis"
	"mentor_app/chatroom/middleware/log"

	"gopkg.in/mgo.v2/bson"
)

func ClassroomInfoMessageProgress(node *dto.WebSocketNode, messageDto *dto.Message, data []byte) error {
	err, getTokenRes := classroom_system_api.GetUserClassroomToken(messageDto.SenderId)
	messageDto.Url = getTokenRes.Data
	err, messageId := saveSingleChatMessage(messageDto)
	if err != nil {
		return err
	}

	messageDto.MessageId = messageId
	messageDto.Time = bson.ObjectIdHex(messageId).Time()

	data, _ = json.Marshal(messageDto)

	_, err = mentor_redis.Client.Publish(chatroom_variables.WS_REDIS_PREFIX+messageDto.SenderId, string(data)).Result()
	if err != nil {
		log.Logger.Error("error publish chat message to sender", err)
	}
	_, err = mentor_redis.Client.Publish(chatroom_variables.WS_REDIS_PREFIX+messageDto.RecieverId, string(data)).Result()
	if err != nil {
		log.Logger.Error("error publish chat message to receiver", err)
	}

	return nil
}
