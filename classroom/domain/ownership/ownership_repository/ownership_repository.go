package ownership_repository

import (
	"context"
	"mentor/classroom/db_connection"
	"mentor/classroom/domain/ownership"
	ownership_dto "mentor/classroom/domain/ownership/dto"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"

	"github.com/qiniu/qmgo"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("classroom").Collection("ownership")
}

func CreateNewOwnership(ctx context.Context, userId string) error {
	newOwnership := ownership.NewOwnership(userId)
	_, err := collection.InsertOne(ctx, newOwnership)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func GetClassroomList(ctx context.Context, userId string) ([]ownership_dto.ClassroomInfo, error) {
	result := ownership_dto.ClassroomList{}
	err := collection.Find(ctx, bson.M{"userId": userId}).Select(bson.M{"_id": 0, "classroomList": 1}).One(&result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result.ClassroomList, nil
}

func GetStudentList(ctx context.Context, userId string) ([]string, error) {
	result := ownership_dto.StudentList{}
	err := collection.Find(ctx, bson.M{"userId": userId}).Select(bson.M{"_id": 0, "studentList": 1}).One(&result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result.StudentList, nil
}

func EnrollNewClassroom(ctx context.Context, userId string, classroomId string, mentorId string) error {
	classroomInfo := ownership.ClassroomInfo{
		MentorId:    mentorId,
		ClassroomId: classroomId,
	}

	err := collection.UpdateOne(ctx, bson.M{"userId": userId}, bson.M{"$pull": bson.M{"classroomList": classroomInfo}})
	if err != nil {
		return errors.WithStack(err)
	}

	err = collection.UpdateOne(ctx, bson.M{"userId": userId}, bson.M{"$push": bson.M{"classroomList": classroomInfo}})
	if err != nil {
		return errors.WithStack(err)
	}

	err = collection.UpdateOne(ctx, bson.M{"userId": mentorId}, bson.M{"$pull": bson.M{"studentList": userId}})
	if err != nil {
		return errors.WithStack(err)
	}

	err = collection.UpdateOne(ctx, bson.M{"userId": mentorId}, bson.M{"$push": bson.M{"studentList": userId}})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
