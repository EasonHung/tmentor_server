package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mentor_app/chatroom/constants/chatroom_variables"
	"mentor_app/chatroom/dao/chat_message_dao"
	"mentor_app/chatroom/dao/conversation_dao"
	"mentor_app/chatroom/mentor_redis"
	"mentor_app/chatroom/vo"

	"mentor_app/chatroom/dao/db_connection"
	"mentor_app/chatroom/dao/ownership_dao"
	"mentor_app/chatroom/dto"
	"mentor_app/chatroom/internalAPI/user_system_api"
	"mentor_app/chatroom/internal_error"
	"mentor_app/chatroom/middleware/log"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func GetUnreadMessageByConversationId(conversationId string, senderId string) (error, []*dto.Message) {
	var err error

	readCursor, err := conversation_dao.FindCursorWithUserId(conversationId, senderId)

	unReadChatMessages := []*chat_message_dao.ChatMessage{}
	err, unReadChatMessages = chat_message_dao.FindByConversationGtMessageId(conversationId, readCursor)
	if err != nil {
		return err, nil
	}

	unReadMessages := []*dto.Message{}
	for _, chatMessage := range unReadChatMessages {
		chatMessage.Data.MessageId = chatMessage.MessageId
		chatMessage.Data.Time = bson.ObjectIdHex(chatMessage.Id).Time()
		unReadMessages = append(unReadMessages, chatMessage.Data)
	}

	return nil, unReadMessages
}

func AddConversation(userId string, participants []string, conversationType int) (string, error) {
	transaction := func(ctx context.Context) (interface{}, error) {
		conversationId := bson.NewObjectId().Hex()
		newConversation := conversation_dao.Conversation{
			ConversationId: conversationId,
			Participants:   participants,
			Type:           conversationType,
			ReadCursor:     make([]conversation_dao.ReadCursor, 0),
		}

		err := ownership_dao.LockWithUserIdList(ctx, participants)
		if err != nil {
			return nil, err
		}

		for _, value := range participants {
			_, err = ownership_dao.FindByUserIdAndNeParticipantsWithTx(ctx, value, getAddConversationTarget(value, participants))
			if err != nil && err.Error() == "mongo: no documents in result" {
				// 無法把中間的結果當成res傳出去 因為會rollback回nil
				return nil, internal_error.DuplicateAddConversationError{}
			} else if err != nil {
				return nil, err
			}

			err = ownership_dao.PushConversationByUserIdWithTx(ctx, value, conversationId, participants, conversationType)
			if err != nil {
				return nil, err
			}
		}

		_, err = conversation_dao.InsertOneWithTx(ctx, newConversation)
		if err != nil {
			return nil, err
		}

		return conversationId, nil
	}

	conversationId, err := db_connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	if err != nil {
		if errors.As(err, &internal_error.DuplicateAddConversationError{}) {
			log.Logger.Info(fmt.Sprintf("duplicated adding conversation, participants:{%s, %s}", participants[0], participants[1]))
			innerErr, conversationDao := conversation_dao.FindByParticipant(participants)
			if innerErr != nil {
				return "", errors.Wrap(err, "")
			}
			return conversationDao.ConversationId, err
		}

		return "", err
	}

	return conversationId.(string), nil
}

func InitUser(userId string) error {
	emptyConversationList := make([]ownership_dao.Conversation, 0)
	ownershipEntity := ownership_dao.Ownership{
		UserId:           userId,
		ConversationList: emptyConversationList,
	}
	if _, err := ownership_dao.InsertOne(&ownershipEntity); err != nil {
		log.Logger.Error(err)
		return err
	}
	return nil
}

func GetUserConversationList(userId string) ([]ownership_dao.Conversation, error) {
	entity, err := ownership_dao.FindByUserId(userId)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	return entity.ConversationList, nil
}

func GetConversationUserInfo(participants []string, userId string) (vo.UserInfoVo, error) {
	userInfoRes := user_system_api.GetUserInfoRes{}
	res := vo.UserInfoVo{}
	var err error

	for _, value := range participants {
		if value != userId {
			err, userInfoRes = user_system_api.GetUserInfo(value)
			if err != nil {
				return res, err
			}

			res.AvatorUrl = userInfoRes.Data.AvatorUrl
			res.Nickname = userInfoRes.Data.Nickname
			res.UserId = userInfoRes.Data.UserId
		}
	}
	return res, nil
}

func GetLastConversationMessageAndTime(conversationId string) (string, time.Time, error) {
	err, entity := chat_message_dao.FindLastByConversationIdAndSortByMessageId(conversationId)

	if err != nil && err.Error() == "mongo: no documents in result" {
		return "", time.Time{}, nil
	}
	if err != nil {
		log.Logger.Error(err)
		return "", time.Time{}, err
	}

	return entity.Data.Message, entity.Data.Time, nil
}

func BatchReadedMsgProcess(userId string, deviceId string, readedMessageIds []string) error {
	readBy := chat_message_dao.ReadedBy{
		UserId:   userId,
		DeviceId: deviceId,
	}
	err := chat_message_dao.UpdateReadedInfoByMessageIdList(readedMessageIds, readBy)
	if err != nil {
		return err
	}

	err, readedMsgs := chat_message_dao.FindByMessageIdList(readedMessageIds)
	if err != nil {
		return errors.Wrap(err, "error get sync message")
	}

	for _, message := range readedMsgs {
		if message.SenderId == userId {
			continue
		}
		data := message.Data
		data.Cmd = 1
		data.MessageId = message.MessageId
		jsonByte, _ := json.Marshal(data)
		_, err := mentor_redis.Client.Publish(chatroom_variables.WS_REDIS_PREFIX+message.SenderId, string(jsonByte)).Result()
		if err != nil {
			log.Logger.Error("error publish readed message to receiver", err)
		}
	}

	return nil
}

func CountUnReadedMessage(conversationId string, userId string) (int64, error) {
	readCursor, err := conversation_dao.FindCursorWithUserId(conversationId, userId)
	if err != nil {
		return 0, err
	}

	err, count := chat_message_dao.CountMessagesAfterId(conversationId, readCursor)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func GetSyncMessages(conversationId string, lastMessageId string) ([]*chat_message_dao.ChatMessage, error) {
	unReadChatMessages := []*chat_message_dao.ChatMessage{}

	if lastMessageId == "" {
		err, lastMessage := chat_message_dao.FindLastByConversationIdAndSortByMessageId(conversationId)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				return unReadChatMessages, nil
			}
			return nil, err
		}
		unReadChatMessages = append(unReadChatMessages, &lastMessage)
		return unReadChatMessages, nil
	}

	err, unReadChatMessages := chat_message_dao.FindByConversationGtMessageId(conversationId, lastMessageId)
	if err != nil {
		return nil, err
	}

	return unReadChatMessages, nil
}

func UpdateReadCursor(conversationId string, userId string, deviceId string, lastMessageId string) error {
	err := conversation_dao.UpsertConversationCursor(conversationId, userId, deviceId, lastMessageId)
	return err
}

func GetAnotherCursor(conversationId string, userId string) (error, string) {
	err, cursors := conversation_dao.GetCursorList(conversationId)
	if err != nil {
		return err, ""
	}

	res := ""
	for _, cursor := range cursors {
		if cursor.UserId == userId {
			continue
		}
		if cursor.Cursor > res {
			res = cursor.Cursor
		}
	}
	return nil, res
}

func GetSelfCursor(conversationId string, userId string) (error, string) {
	err, cursors := conversation_dao.GetCursorList(conversationId)
	if err != nil {
		return err, ""
	}

	res := ""
	for _, cursor := range cursors {
		if cursor.UserId != userId {
			continue
		}
		if cursor.Cursor > res {
			res = cursor.Cursor
		}
	}
	return nil, res
}

func getAddConversationTarget(adder string, participants []string) string {
	for _, user := range participants {
		if user != adder {
			return user
		}
	}
	return ""
}

func changeCmdByReadBy(userId string, messages []chat_message_dao.ChatMessage) []chat_message_dao.ChatMessage {
	for index, message := range messages {
		readed := false
		for _, readedBy := range message.ReadedBy {
			if readedBy.UserId == userId {
				readed = true
				break
			}
		}

		if readed {
			messages[index].Data.Cmd = 1
		}
	}

	return messages
}
