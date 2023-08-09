package ownership_dao

import (
	"context"
	"fmt"
	"mentor_app/chatroom/dao/db_connection"

	"github.com/pkg/errors"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("chatroom").Collection("ownership")
}

func InsertOne(obj *Ownership) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.Background(), obj)
	return result, err
}

func PushConversationByUserIdWithTx(ctx context.Context, userId string, conversationId string, participantsId []string, conversationType int) error {
	conversationItem := Conversation{
		ConversationId: conversationId,
		Participants:   participantsId,
		Type:           conversationType,
	}

	err := collection.UpdateOne(ctx,
		bson.M{"userId": userId},
		bson.M{"$push": bson.M{
			"conversationList": conversationItem,
		}})
	return err
}

func FindByUserId(userId string) (Ownership, error) {
	result := Ownership{}
	err := collection.Find(context.Background(), bson.M{"userId": userId}).One(&result)
	if err != nil {
		fmt.Println("here")
	}
	return result, err
}

func FindByUserIdAndParticipantsWithTx(ctx context.Context, userId string, participant string) (Ownership, error) {
	result := Ownership{}
	err := collection.Find(context.Background(), bson.M{"userId": userId, "conversationList.participants": participant}).One(&result)
	if err != nil {
		return result, errors.Wrap(err, "error get ownership")
	}
	return result, err
}

func FindByUserIdAndNeParticipantsWithTx(ctx context.Context, userId string, participants string) (Ownership, error) {
	result := Ownership{}
	err := collection.Find(ctx, bson.M{"$or": []interface{}{
		bson.M{"$and": []interface{}{bson.M{"userId": userId, "conversationList": []interface{}{}}}},
		bson.M{"$and": []interface{}{bson.M{"userId": userId, "conversationList.participants": bson.M{"$ne": participants}}}},
	},
	}).One(&result)
	if err != nil {
		return result, err
	}
	return result, err
}

func LockWithUserIdList(ctx context.Context, userIdList []string) error {
	_, err := collection.UpdateAll(ctx, bson.M{"userId": bson.M{"$in": userIdList}}, bson.M{"$set": bson.M{"Lock": primitive.NewObjectID()}})
	if err != nil {
		return errors.Wrap(err, "error lock owner")
	}
	return err
}
