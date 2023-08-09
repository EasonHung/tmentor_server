package user_evaluation_dao

import (
	"context"
	"evaluation_system/dao/db_connection"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("evaluation_system").Collection("user_evaluation")
}

func CreateOne(obj UserEvaluation) (*qmgo.InsertOneResult, error) {
	obj.CreateTime = time.Now()
	obj.ModifyTime = time.Now()
	result, err := collection.InsertOne(context.Background(), obj)
	if err != nil {
		err = errors.Wrap(err, "error create user evaluation")
		return nil, err
	}
	return result, err
}

func PushPostIdByUserIdWithTx(ctx context.Context, userId string, postId string) error {
	err := collection.UpdateOne(ctx, bson.M{"userId": userId}, bson.M{"$push": bson.M{"evaluationPostIds": postId}})
	return err
}

func FindByUserId(userId string) (UserEvaluation, error) {
	var userEvaluation UserEvaluation
	err := collection.Find(context.Background(), bson.M{"userId": userId}).One(&userEvaluation)
	if err != nil {
		err = errors.Wrap(err, "error find user evaluation")
		return userEvaluation, err
	}

	return userEvaluation, nil
}

func FindByUserIdWithTx(ctx context.Context, userId string) (UserEvaluation, error) {
	var userEvaluation UserEvaluation
	err := collection.Find(ctx, bson.M{"userId": userId}).One(&userEvaluation)
	if err != nil {
		err = errors.Wrap(err, "error find user evaluation")
		return userEvaluation, err
	}

	return userEvaluation, nil
}

func UpdateAverageScoreByUserIdWithTx(ctx context.Context, userId string, averageScore float64) error {
	err := collection.UpdateOne(ctx, bson.M{"userId": userId},
		bson.M{
			"$set": bson.M{
				"averageScore": averageScore,
			},
		})
	if err != nil {
		return err
	}

	return nil
}
