package conversation_dao

import (
	"context"
	"mentor_app/chatroom/dao/db_connection"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("chatroom").Collection("conversation")
}

func InsertOne(obj Conversation) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.Background(), obj)
	return result, err
}

func InsertOneWithTx(ctx context.Context, obj Conversation) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, obj)
	return result, err
}

func FindByConversationId(conversationId string) (error, []string) {
	conversation := Conversation{}
	err := collection.Find(context.Background(), bson.M{"_id": conversationId}).One(&conversation)
	if err != nil {
		return err, nil
	}

	memberIds := []string{}
	for _, memberId := range conversation.Participants {
		memberIds = append(memberIds, memberId)
	}

	return nil, memberIds
}

func FindByParticipant(participants []string) (error, Conversation) {
	conversation := Conversation{}
	err := collection.Find(context.Background(), bson.M{"participants": bson.M{ "$size": 2, "$all": participants}}).One(&conversation)
	if err != nil {
		return err, conversation
	}
	return nil, conversation
}

func GetCursorList(conversationId string) (error, []ReadCursor) {
	conversation := Conversation{}
	err := collection.Find(context.TODO(), bson.M{"conversationId": conversationId}).One(&conversation)
	if err != nil {
		return errors.WithStack(err), nil
	}
	return nil, conversation.ReadCursor
}

func UpsertConversationCursor(conversationId string, userId string, deviceId string, messageId string) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		err := collection.UpdateOne(ctx,
			bson.M{"conversationId": conversationId},
			bson.M{"$pull": bson.M{"readCursor": bson.M{"userId": userId, "deviceId": deviceId}}},
		)
		if err != nil {
			return nil, err
		}

		newCursor := ReadCursor{
			UserId:   userId,
			DeviceId: deviceId,
			Cursor:   messageId,
		}
		err = collection.UpdateOne(ctx,
			bson.M{"conversationId": conversationId},
			bson.M{"$push": bson.M{"readCursor": newCursor}},
		)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err := db_connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	return err
}

func FindCursor(conversationId string, userId string, deviceId string) (string, error) {
	conversation := Conversation{}
	err := collection.Find(context.Background(),
		bson.M{"conversationId": conversationId}).One(&conversation)
	if err != nil {
		return "", err
	}

	for _, cursor := range conversation.ReadCursor {
		if cursor.UserId == userId && cursor.DeviceId == deviceId {
			return cursor.Cursor, nil
		}
	}

	return "", qmgo.ErrNoSuchDocuments
}

func FindCursorWithUserId(conversationId string, userId string) (string, error) {
	conversation := Conversation{}
	err := collection.Find(context.Background(),
		bson.M{"conversationId": conversationId}).One(&conversation)
	if err != nil {
		return "", err
	}

	readCursor := ""
	for _, cursor := range conversation.ReadCursor {
		if cursor.UserId == userId && cursor.Cursor > readCursor {
			readCursor = cursor.Cursor
		}
	}

	if readCursor == "" {
		return "", qmgo.ErrNoSuchDocuments
	}
	return readCursor, nil
}
