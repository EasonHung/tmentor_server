package message_service

import (
	"encoding/json"
	"mentor/classroom/constants/application_type"
	"mentor/classroom/constants/redis_prefix"
	"mentor/classroom/domain/classroom_status/classroom_status_repository"
	"mentor/classroom/mentor_redis"
	"mentor/classroom/service/list_service"
	"strings"

	"github.com/pkg/errors"
)

func SendMessage(applicationType string, receiverId string, message any) error {
	jsonByte, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}

	err = mentor_redis.Client.Publish(redis_prefix.WS_REDIS_PREFIX+applicationType+"-"+receiverId, string(jsonByte)).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func SendMessageToAllApplicationType(receiverId string, message any) error {
	jsonByte, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}

	err = mentor_redis.Client.Publish(redis_prefix.WS_REDIS_PREFIX+application_type.APP+"-"+receiverId, string(jsonByte)).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	err = mentor_redis.Client.Publish(redis_prefix.WS_REDIS_PREFIX+application_type.WEB+"-"+receiverId, string(jsonByte)).Err()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func SendMessageToSpecificClassroom(classroomId string, message any) error {
	jsonByte, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}

	memberList, err := classroom_status_repository.GetClassroomMembers(classroomId)
	if err != nil {
		return err
	}
	for _, userIdWithApplicationType := range memberList {
		err = mentor_redis.Client.Publish(redis_prefix.WS_REDIS_PREFIX+userIdWithApplicationType, string(jsonByte)).Err()
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func SendMessageToOtherMembersInClassroom(classroomId string, userId string, message any) error {
	jsonByte, err := json.Marshal(message)
	if err != nil {
		return errors.WithStack(err)
	}

	memberList, err := classroom_status_repository.GetClassroomMembers(classroomId)
	if err != nil {
		return err
	}
	for _, userIdWithApplicationType := range memberList {
		userIdInClassroom := strings.Split(userIdWithApplicationType, "-")[1]
		if(userId == userIdInClassroom) {
			continue
		}

		err = mentor_redis.Client.Publish(redis_prefix.WS_REDIS_PREFIX+userIdWithApplicationType, string(jsonByte)).Err()
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func SendMessageToAllStudentAndMentor(classroomId string, message any) error {
	mentorId, studentList, err := list_service.GetMentorIdAndStudentListByClassroomId(classroomId)
	if err != nil {
		return err
	}

	SendMessageToAllApplicationType(mentorId, message)
	for _, studentId := range studentList {
		err = SendMessageToAllApplicationType(studentId, message)
		if err != nil {
			return err
		}
	}

	return nil
}
