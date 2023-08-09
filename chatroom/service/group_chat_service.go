package service

import (
	"encoding/json"
	"mentor_app/chatroom/dao/chat_message_dao"
	"mentor_app/chatroom/dao/conversation_dao"
	"mentor_app/chatroom/dto"
	"mentor_app/chatroom/middleware/log"

	"gopkg.in/mgo.v2/bson"
)

func GroupChatMessageProgress(node *dto.WebSocketNode, messageDto *dto.Message, data []byte) error {
	err, messageId := groupChatMessageSaveProgress(messageDto)
	if err != nil {
		return err
	}

	node.DataQueue <- data
	recieverIdList := []string{}
	recieverIdList, hasCachedList := conversationMemberMap[messageDto.ConversationId]
	if !hasCachedList {
		_, recieverIdList = conversation_dao.FindByConversationId(messageDto.ConversationId)
		conversationMemberMap[messageDto.ConversationId] = recieverIdList
	}

	for _, recieverId := range recieverIdList {
		if recieverId == messageDto.SenderId {
			continue
		}

		rwLock.Lock()
		recieverNode, hasConnected := SocketConnectionMap[recieverId]
		rwLock.Unlock()
		if hasConnected {
			messageDto.MessageId = messageId
			messageDto.Time = bson.ObjectIdHex(messageId).Time()
			data, _ = json.Marshal(messageDto)
			recieverNode.DataQueue <- data
		}
	}

	return nil
}

func GroupChatReadedProgress(messageDto *dto.Message, data []byte) error {
	readBy := chat_message_dao.ReadedBy{
		UserId: messageDto.SenderId,
	}
	if err := chat_message_dao.UpdateReadedInfoByMessageId(messageDto.MessageId, readBy); err != nil {
		log.Logger.Error("[DB] error occured when update readed message", err)
		return err
	}

	recieverIdList := []string{}
	recieverIdList, hasCachedList := conversationMemberMap[messageDto.ConversationId]
	if !hasCachedList {
		_, recieverIdList = conversation_dao.FindByConversationId(messageDto.ConversationId)
		conversationMemberMap[messageDto.ConversationId] = recieverIdList
	}

	for _, recieverId := range recieverIdList {
		if recieverId == messageDto.SenderId {
			continue
		}

		rwLock.Lock()
		recieverNode, ok := SocketConnectionMap[recieverId]
		rwLock.Unlock()
		if ok {
			recieverNode.DataQueue <- data
		}
	}

	return nil
}

// todo: add transaction
func groupChatMessageSaveProgress(dto *dto.Message) (error, string) {
	// save message
	messageId := bson.NewObjectId().Hex()
	message := chat_message_dao.ChatMessage{
		Id:             messageId,
		SenderId:       dto.SenderId,
		Data:           dto,
		ReadedBy:       []chat_message_dao.ReadedBy{},
		ConversationId: dto.ConversationId,
	}
	_, err := chat_message_dao.InsertOne(&message)
	if err != nil {
		log.Logger.Error("[DB] error occured when insert message", err)
		return err, ""
	}
	return err, messageId
}
