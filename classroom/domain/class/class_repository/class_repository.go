package class_repository

import (
	"context"
	"mentor/classroom/db_connection"
	"mentor/classroom/domain/class"
	"mentor/classroom/domain/class/constants/class_status"
	class_dto "mentor/classroom/domain/class/dto"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("classroom").Collection("class")
}

func FindLastClass(ctx context.Context, classroomId string, studentId string) (error, class_dto.LastClassInfo) {
	result := class_dto.LastClassInfo{}
	err := collection.Find(ctx, bson.M{"classroomId": classroomId, "studentId": studentId}).Sort("-classId").One(&result)
	if err != nil {
		return errors.WithStack(err), result
	}

	return nil, result
}

func InitClass(ctx context.Context, classroomId string, mentorId string, studentId string, classTitle string, classDesc string, points int, classTime int) (error, string) {
	newClass := class.NewClass(classroomId, mentorId, studentId, classTitle, classDesc, points, classTime)
	_, err := collection.InsertOne(ctx, newClass)
	if err != nil {
		return errors.WithStack(err), ""
	}
	return nil, newClass.ClassId
}

func ReimburseClass(ctx context.Context, classId string, classTime int) error {
	err := collection.UpdateOne(ctx, bson.M{"classId": classId}, bson.M{"$set": bson.M{"status": class_status.Reimburse, "classTime": classTime, "remainTime": classTime}})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func StartClass(ctx context.Context, classId string) error {
	err := collection.UpdateOne(ctx, bson.M{"classId": classId}, bson.M{"$set": bson.M{"status": class_status.Start, "startTime": time.Now()}})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func FinishClass(ctx context.Context, classId string) error {
	err := collection.UpdateOne(ctx, bson.M{"classId": classId}, bson.M{"$set": bson.M{"status": class_status.Finish}})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func ClockOn(ctx context.Context, classId string, userId string, newRemainTime int) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		err := collection.UpdateOne(ctx, bson.M{"classId": classId}, bson.M{"$pull": bson.M{"clockRecord": bson.M{"userId": userId}}})
		if err != nil {
			return nil, errors.WithStack(err)
		}

		newRecord := class.ClockRecord{
			UserId:    userId,
			ClockTime: time.Now(),
		}
		err = collection.UpdateOne(ctx, bson.M{"classId": classId}, bson.M{"$push": bson.M{"clockRecord": newRecord}})
		if err != nil {
			return nil, errors.WithStack(err)
		}

		err = collection.UpdateOne(ctx, bson.M{"classId": classId}, bson.M{"$set": bson.M{"remainTime": newRemainTime}})
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return nil, nil
	}

	_, err := db_connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	return err
}

func FindByClassId(ctx context.Context, classId string) (error, class_dto.ClassInfo) {
	result := class_dto.ClassInfo{}
	err := collection.Find(ctx, bson.M{"classId": classId}).One(&result)
	if err != nil {
		return errors.WithStack(err), result
	}

	return nil, result
}

func FindClassRecord(ctx context.Context, userId string) (error, []class_dto.ClassRecord) {
	result := make([]class_dto.ClassRecord, 0)
	err := collection.Find(ctx, bson.M{
		"$and": []bson.M{
			{
				"status": bson.M{"$ne": class_status.Init},
			},
			{
				"$or": []bson.M{
					{"mentorId": userId},
					{"studentId": userId},
				},
			},
		},
	}).All(&result)
	if err != nil {
		return errors.WithStack(err), result
	}

	return nil, result
}
