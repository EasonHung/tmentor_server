package class_info_dao

import (
	"context"
	"mentor/classroom/dao/class_info_dao/enum/class_info_status"
	"mentor/classroom/dao/db_connection"
	"time"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("classroom").Collection("class_info")
}

func InsertOne(obj *ClassInfo) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(context.Background(), obj)
	return result, err
}

func FindLastByClassroomIdStudentIdAndStatus(classroomId string, studentUserId string, status class_info_status.Status) (ClassInfo, error) {
	result := ClassInfo{}
	err := collection.Find(context.Background(), bson.M{"classroomId": classroomId, "studentUserId": studentUserId, "status": status.String()}).Sort("-createTime").One(&result)
	return result, err
}

func FindByClassId(classId string) (ClassInfo, error) {
	result := ClassInfo{}
	classObjectId, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return result, err
	}
	err = collection.Find(context.Background(), bson.M{"_id": classObjectId}).One(&result)
	return result, err
}

func UpdateStatusByClassId(classId string, status class_info_status.Status) error {
	classObjectId, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return err
	}
	err = collection.UpdateOne(context.Background(), bson.M{"_id": classObjectId, "status": bson.M{"$ne": status.String()}}, bson.M{"$set": bson.M{"status": status.String(), "modifyTime": time.Now()}})
	return err
}

func UpdateStartTimeByClassId(classId string, startTime time.Time) error {
	classObjectId, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return err
	}
	err = collection.UpdateOne(context.Background(), bson.M{"_id": classObjectId}, bson.M{"$set": bson.M{"startTime": startTime, "modifyTime": time.Now()}})
	return err
}

func UpdateRemainTimeByClassId(classId string, remainTime int) error {
	classObjectId, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return err
	}
	err = collection.UpdateOne(context.Background(), bson.M{"_id": classObjectId}, bson.M{"$set": bson.M{"remainTime": remainTime, "modifyTime": time.Now()}})
	return err
}
