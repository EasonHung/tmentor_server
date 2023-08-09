package student_case_dao

import (
	"case_system/dao/db_connection"
	"case_system/dto"
	"context"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("case_system").Collection("student_case")
}

func InsertOne(obj StudentCase) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.Background(), obj)
	return result, err
}

func FindWithPagination(page int64, pagePerSize int64) ([]StudentCase, error) {
	result := []StudentCase{}
	err := collection.Find(context.Background(), bson.M{}).Skip(page * pagePerSize).Limit(pagePerSize).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func PushBidInfoByCaseIdWithTx(ctx context.Context, studentCaseId string, bidInfoId string) error {
	err := collection.UpdateOne(ctx,
		bson.M{"studentCaseId": studentCaseId},
		bson.M{"$push": bson.M{
			"bidInfoIds": bidInfoId,
		}})
	return err
}

func FindBidInfoIdByCaseId(studentCaseId string) ([]string, error) {
	var result dto.BidInfoIds

	err := collection.Find(context.Background(), bson.M{"studentCaseId": studentCaseId}).Select(bson.M{"_id": 0, "bidInfoIds": 1}).One(&result)
	if err != nil {
		return nil, err
	}
	return result.BidInfoIds, nil
}

func FindCasesByUserIdDesc(userId string) ([]StudentCase, error) {
	result := []StudentCase{}
	err := collection.Find(context.Background(), bson.M{"userId": userId}).Sort("-postTime").All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindByCaseId(studentCaseId string) (StudentCase, error) {
	result := StudentCase{}

	err := collection.Find(context.Background(), bson.M{"studentCaseId": studentCaseId}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
