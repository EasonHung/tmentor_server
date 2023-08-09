package classroom_status_repository

import (
	"mentor/classroom/constants/application_type"
	"mentor/classroom/constants/classroom_status_enum"
	"mentor/classroom/constants/redis_prefix"
	"mentor/classroom/dto/ws_message"
	"mentor/classroom/mentor_redis"
	"mentor/classroom/utils/string_utils"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

func GetClassroomStatus(classroomId string, mentorId string) (string, error) {
	peopleInClassroom, err := mentor_redis.Client.SMembers(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX + classroomId).Result()
	if err != nil {
		return "", errors.New(err.Error())
	}

	return getClassroomStatus(peopleInClassroom, mentorId), nil
}

// todo: use redis transaction
func OpenClassroom(classroomId string, userAppType string, userId string, classSettingInfo ws_message.ClassSettingInfo) error {
	// check if classroom is already opened
	peopleInClassroom, err := mentor_redis.Client.SMembers(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX + classroomId).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	if len(peopleInClassroom) != 0 {
		return errors.New("classroom is already opened!")
	}

	// push member into class
	res, err := mentor_redis.Client.SAdd(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX+classroomId, userAppType+"-"+userId).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	// check if add successed
	if res != 1 {
		return errors.New("open classroom failed, for add to set failed!")
	}

	// set class setting into redis
	classSettingInfo.ClassroomId = classroomId
	_, classSettingInfoJsonString := classSettingInfo.ToJsonString()
	resStr, err := mentor_redis.Client.Set(redis_prefix.CLASS_INFO_REDIS_PREFIX+classroomId, classSettingInfoJsonString, 0).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	if resStr != "OK" {
		return errors.New("open classroom failed, for set classroom information failed!, res string: " + resStr)
	}
	return nil
}

func JoinClassroom(classroomId string, applicationType string, userId string) (string, error) {
	joinRoomScript := redis.NewScript(`
		local memberList = redis.call("smembers", KEYS[1])

		if( #memberList == 0) then
			return "class not open"
		elseif( #memberList > 2 ) then
			return "already in class"
		else
			redis.call('sadd', KEYS[1], KEYS[2])
			return "join class success"
		end

		return "join fail"
	`)
	res, err := joinRoomScript.Run(mentor_redis.Client, []string{redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX + classroomId, applicationType + "-" + userId}).Result()
	// _, err = mentor_redis.Client.LPush(class_channel_var.OPENED_CLASSMEMBER_LIST_PREFIX+classRoomId, dto.ApplicationType+"-"+dto.SenderId).Result()
	if err != nil {
		return "", errors.New("join room fail")
	}
	return res.(string), nil
}

func CloseClassroom(classroomId string) error {
	err := mentor_redis.Client.Del(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX + classroomId).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func LeaveClassroom(classroomId string, studentApplicationType string, studentId string) error {
	err := mentor_redis.Client.SRem(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX+classroomId, studentApplicationType+"-"+studentId).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func GetClassroomMembers(classroomId string) ([]string, error) {
	memberList, err := mentor_redis.Client.SMembers(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX + classroomId).Result()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return memberList, nil
}

func getClassroomStatus(peopleInClassroom []string, userId string) string {
	if string_utils.ContainString(peopleInClassroom, application_type.WEB+"-"+userId) ||
		string_utils.ContainString(peopleInClassroom, application_type.APP+"-"+userId) {
		if len(peopleInClassroom) > 1 {
			return classroom_status_enum.IN_CLASS
		} else {
			return classroom_status_enum.ONLINE
		}
	} else {
		return classroom_status_enum.OFFLINE
	}
}
