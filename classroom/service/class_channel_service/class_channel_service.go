package class_channel_service

import (
	"mentor/classroom/constants/redis_prefix"
	"mentor/classroom/domain/classroom/classroom_repository"
	"mentor/classroom/domain/classroom_status/classroom_status_repository"
	"mentor/classroom/dto"
	"mentor/classroom/mentor_redis"
	"mentor/classroom/middleware/log"
	"mentor/classroom/service/list_service"
	"mentor/classroom/service/message_service"

	"github.com/pkg/errors"
)

// applicationType是有"-"的
func NotifyClosureAndCleanMap(userId string, applicationType string, locatedClassroomId string) error {
	if locatedClassroomId == "" {
		return nil
	}

	mentorId, studentList, err := list_service.GetMentorIdAndStudentListByClassroomId(locatedClassroomId)
	if err != nil {
		return err
	}

	if mentorId == userId {
		// close whole classroom
		err := classroom_status_repository.CloseClassroom(locatedClassroomId)
		if err != nil {
			return err
		}

		err = classroom_repository.DeactiveClassSetting(locatedClassroomId)
		if err != nil {
			return err
		}

		// inform everyone member who cares classroom status, including mentor himself (mentor may be using another device to watch classroom status).
		message := dto.GenerateCloseClassroomCmd(locatedClassroomId, userId, applicationType)
		for _, studentId := range studentList {
			err = message_service.SendMessageToAllApplicationType(studentId, message)
			if err != nil {
				return err
			}
		}
	} else {
		// leave classroom
		message := dto.GenerateLeaveClassroomCmd(locatedClassroomId, userId, applicationType)
		err := message_service.SendMessageToSpecificClassroom(locatedClassroomId, message)
		if err != nil {
			return err
		}

		err = classroom_status_repository.LeaveClassroom(locatedClassroomId, applicationType, userId)
		if err != nil {
			return err
		}

		err = message_service.SendMessageToAllStudentAndMentor(locatedClassroomId, message)
		if err != nil {
			return errors.New("send leave classroom message fail!")
		}
	}
	// leaveMsg := "leave"
	// classroomId := class_channel_var.UserLocateMap[userId]

	// memberList, err := mentor_redis.Client.SMembers(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX + classroomId).Result()
	// if err != nil {
	// 	return errors.Wrap(err, "error find memberlist")
	// }

	// for _, memberId := range memberList {
	// 	if memberId == userId {
	// 		continue
	// 	}

	// 	message := dto.Message{
	// 		Cmd: leaveMsg,
	// 	}
	// 	data, err := json.Marshal(message)
	// 	if err != nil {
	// 		return errors.Wrap(err, "error notify member due to json marshal")
	// 	}

	// 	mentor_redis.Client.Publish(memberId, string(data))
	// }

	// if len(memberList) > 0 && len(memberList) < 2 {
	// 	mentor_redis.Client.Del(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX + classroomId)
	// 	studentList, err := list_service.GetStudentListByClassroomId(classroomId)
	// 	if err != nil {
	// 		log.Logger.Error("load classroom info fail, ", err.Error())
	// 		return err
	// 	}
	// 	for _, studentId := range studentList {
	// 		message := dto.Message{
	// 			Cmd:     "offline",
	// 			Message: classroomId,
	// 		}
	// 		jsonByte, _ := json.Marshal(message)
	// 		mentor_redis.Client.Publish(class_channel_var.WS_REDIS_PREFIX+class_channel_var.WEB+studentId, string(jsonByte))
	// 		mentor_redis.Client.Publish(class_channel_var.WS_REDIS_PREFIX+class_channel_var.APP+studentId, string(jsonByte))
	// 	}
	// } else {
	// 	mentor_redis.Client.SRem(redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX+classroomId, applicationType+userId)
	// }

	// err = classroom_repository.DeactiveClassSetting(classroomId)
	// if err != nil {
	// 	log.Logger.Error("load classroom info fail, ", err.Error())
	// 	return err
	// }

	// delete(class_channel_var.UserLocateMap, userId)
	return nil
}

func GetClassroomList() []string {
	keys, _, err := mentor_redis.Client.Scan(0, redis_prefix.OPENED_CLASSMEMBER_LIST_PREFIX+"*", 1000).Result()
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		return nil
	}
	return keys
}
