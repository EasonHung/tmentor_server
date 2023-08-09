package user_info_dao

import (
	"context"
	"mentor_app/chatroom/dao/db_connection"
	"mentor_app/chatroom/dto"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("user_info_system").Collection("user_info")
}

func FindFcmTokenByUserId(userId string) (dto.UserFcmTokenNickname, error) {
	result := dto.UserFcmTokenNickname{}
	err := collection.Find(context.Background(), bson.M{"userId": userId}).Select(bson.M{"fcmToken": 1, "nickname": 1}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func FindAvatarUrlByUserId(userId string) (string, error) {
	result := dto.UserAvator{}
	err := collection.Find(context.Background(), bson.M{"userId": userId}).Select(bson.M{"avatorUrl": 1}).One(&result)
	if err != nil {
		return "", err
	}

	return result.AvatorUrl, nil
}
