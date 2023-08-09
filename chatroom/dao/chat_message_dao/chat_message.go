package chat_message_dao

import (
	"mentor_app/chatroom/dto"
)

type ChatMessage struct {
	Id             string       `bson:"_id,omitempty"`
	MessageId      string       `bson:"messageId"`
	SenderId       string       `bson:"senderId"`
	Data           *dto.Message `bson:"data"`
	ConversationId string       `bson:"conversationId"`
	ReadCount      int          `bson:"readCount"`
	ReadedBy       []ReadedBy   `bson:"readedBy"`
}

type ReadedBy struct {
	DeviceId string `bson:"deviceId"`
	UserId   string `bson:"userId"`
}
