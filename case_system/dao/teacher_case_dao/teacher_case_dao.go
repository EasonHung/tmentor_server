package teacher_case_dao

import (
	"case_system/dao/db_connection"
	"context"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("case_system").Collection("teacher_case")
}

func InsertOne(obj TeacherCase) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.Background(), obj)
	return result, err
}

func FindWithPagination(page int64, pagePerSize int64) ([]TeacherCase, error) {
	result := []TeacherCase{}
	err := collection.Find(context.Background(), bson.M{}).Skip(page * pagePerSize).Limit(pagePerSize).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
