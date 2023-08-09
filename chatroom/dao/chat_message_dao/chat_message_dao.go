package chat_message_dao

import (
	"context"
	"mentor_app/chatroom/dao/db_connection"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("chatroom").Collection("message")
}

func InsertOne(obj *ChatMessage) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.Background(), obj)
	return result, err
}

func FindByConversationIdAndSortByMessageId(conversationId string) (error, []ChatMessage) {
	chatMessages := []ChatMessage{}

	err := collection.Find(context.Background(), bson.M{"conversationId": conversationId}).Sort("messageId").All(&chatMessages)
	if err != nil {
		return err, nil
	}

	return nil, chatMessages
}

func FindLastByConversationIdAndSortByMessageId(conversationId string) (error, ChatMessage) {
	chatMessage := ChatMessage{}

	err := collection.Find(context.Background(), bson.M{"conversationId": conversationId}).Sort("-messageId").Limit(1).One(&chatMessage)
	if err != nil {
		return errors.WithStack(err), chatMessage
	}

	return nil, chatMessage
}

func UpdateReadedInfoByMessageId(messageId string, readedBy ReadedBy) error {
	err := collection.UpdateOne(context.Background(), bson.M{"messageId": messageId}, bson.M{"$push": bson.M{"readedBy": readedBy}, "$inc": bson.M{"readCount": 1}})
	if err != nil {
		return err
	}

	return nil
}

func UpdateReadedInfoByMessageIdList(messageIds []string, readedBy ReadedBy) error {
	_, err := collection.UpdateAll(context.Background(), bson.M{"messageId": bson.M{"$in": messageIds}}, bson.M{"$push": bson.M{"readedBy": readedBy}, "$inc": bson.M{"readCount": 1}})
	if err != nil {
		return err
	}

	return nil
}

func CountUnReadedMessage(conversationId string, readCount int, senderId string) (error, int64) {
	count, err := collection.Find(context.Background(), bson.M{"conversationId": conversationId, "readCount": readCount, "senderId": bson.M{"$ne": senderId}}).Count()
	if err != nil {
		return err, -1
	}

	return nil, count
}

func FindByMessageIdList(messageIds []string) (error, []ChatMessage) {
	batch := []ChatMessage{}
	err := collection.Find(context.Background(), bson.M{"messageId": bson.M{"$in": messageIds}}).All(&batch)
	if err != nil {
		return err, nil
	}

	return nil, batch
}

func CountMessagesAfterId(conversationId string, messageId string) (error, int64) {
	count, err := collection.Find(context.Background(), bson.M{"conversationId": conversationId, "messageId": bson.M{"$gt": messageId}}).Count()
	if err != nil {
		return err, 0
	}
	return nil, count
}

func FindByConversationGtMessageId(conversationId string, messageId string) (error, []*ChatMessage) {
	chatMessages := []*ChatMessage{}
	err := collection.Find(context.Background(), bson.M{"conversationId": conversationId, "messageId": bson.M{"$gt": messageId}}).All(&chatMessages)
	if err != nil {
		return err, nil
	}

	return nil, chatMessages
}
