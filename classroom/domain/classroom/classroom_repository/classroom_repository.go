package classroom_repository

import (
	"context"
	"mentor/classroom/constants/redis_prefix"
	"mentor/classroom/db_connection"
	"mentor/classroom/domain/classroom"
	"mentor/classroom/domain/classroom/dto"
	"mentor/classroom/mentor_redis"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("classroom").Collection("classroom")
}

func FindClassroomIdByMentorId(ctx context.Context, mentorId string) (string, error) {
	result := classroom_dto.ClassroomId{}
	err := collection.Find(ctx, bson.M{"mentorId": mentorId}).Select(bson.M{"_id": 0, "classroomId": 1}).One(&result)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return result.ClassroomId, nil
}

func FindMentorIdByClassroomId(ctx context.Context, classroomId string) (string, error) {
	result := classroom_dto.MentorId{}
	err := collection.Find(ctx, bson.M{"classroomId": classroomId}).Select(bson.M{"_id": 0, "mentorId": 1}).One(&result)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return result.MentorId, nil
}

func CreateNewClassroom(ctx context.Context, mentorId string) error {
	newClassroom := classroom.NewClassroom(mentorId)
	_, err := collection.InsertOne(ctx, newClassroom)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func GetActiveClassSetting(classroomId string) (classroom_dto.ClassSetting, error) {
	classSettingInfo := classroom_dto.ClassSetting{}
	// when user open class, system will save its current setting to redis. 
	// key is redis_prefix.CLASS_INFO_REDIS_PREFIX + classroomId.
	classInfoJsonString, err := mentor_redis.Client.Get(redis_prefix.CLASS_INFO_REDIS_PREFIX + classroomId).Result()
	if err != nil {
		return classSettingInfo, errors.WithStack(err)
	}

	classSettingInfo.ParseFromJsonString(classInfoJsonString)
	return classSettingInfo, nil
}

func DeactiveClassSetting(classroomId string) error {
	err := mentor_redis.Client.Del(redis_prefix.CLASS_INFO_REDIS_PREFIX + classroomId).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func GetStoredClassSetting(ctx context.Context, mentorId string) ([]classroom_dto.ClassSetting, error) {
	result := classroom_dto.ClassSettingList{}
	err := collection.Find(ctx, bson.M{"mentorId": mentorId}).Select(bson.M{"_id": 0, "classSettingList": 1}).One(&result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result.ClassSettingList, nil
}

func AddNewClassSetting(ctx context.Context, mentorId string, newClassSetting classroom.ClassSetting) error {
	err := collection.UpdateOne(ctx, bson.M{"mentorId": mentorId}, bson.M{"$push": bson.M{"classSettingList": newClassSetting}})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// if there is no same setting name at first, then this function will just like AddNewClassSetting
func UpdateClassSetting(ctx context.Context, mentorId string, updatedClassSetting classroom.ClassSetting) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		err := collection.UpdateOne(ctx,
			bson.M{"mentorId": mentorId},
			bson.M{"$pull": bson.M{"classSettingList": bson.M{"settingName": updatedClassSetting.SettingName}}},
		)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		err = collection.UpdateOne(ctx,
			bson.M{"mentorId": mentorId},
			bson.M{"$push": bson.M{"classSettingList": updatedClassSetting}},
		)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err := db_connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	return err
}

func DeleteClassSetting(ctx context.Context, mentorId string, classSettingName string) error {
	err := collection.UpdateOne(ctx, bson.M{"mentorId": mentorId}, bson.M{"$pull": bson.M{"classSettingList": bson.M{"settingName": classSettingName}}})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
