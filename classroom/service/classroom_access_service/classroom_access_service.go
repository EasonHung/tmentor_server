package classroom_access_service

import (
	"context"
	"fmt"
	"mentor/classroom/constants/application_type"
	"mentor/classroom/constants/ws_cmd"
	"mentor/classroom/domain/classroom/classroom_repository"
	"mentor/classroom/domain/classroom_status/classroom_status_repository"
	"mentor/classroom/dto"
	"mentor/classroom/dto/ws_message"
	"mentor/classroom/middleware/log"
	"mentor/classroom/service/classroom_registry_service"
	"mentor/classroom/service/jwt_service"
	"mentor/classroom/service/list_service"
	"mentor/classroom/service/message_service"
	"mentor/classroom/service/user_info_service"

	"github.com/pkg/errors"
)

// todo: close class room -> make sure every socket is close in room member

func OpenClassroom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) error {
	// open classroom
	err, classSettingInfo := ws_message.ParseFromJsonString(senderMessage.Message.(string))
	if err != nil {
		return err
	}
	err = classroom_status_repository.OpenClassroom(senderMessage.ClassroomId, senderMessage.ApplicationType, senderMessage.SenderId, classSettingInfo)
	if err != nil {
		return err
	}

	// locate user
	senderWsNode.LocatedClassroomId = senderMessage.ClassroomId
	// class_channel_var.UserLocateMap[senderMessage.SenderId] = senderMessage.ClassroomId // todo: remove this?

	// return success message
	successMessage := senderMessage.FromOpenClassroomCmdToSuccessCmd()
	err = message_service.SendMessageToAllApplicationType(senderMessage.SenderId, successMessage)
	if err != nil {
		return err
	}

	// notify students
	studentList, err := list_service.GetStudentListByClassroomId(senderMessage.ClassroomId)
	if err != nil {
		return err
	}
	for _, studentId := range studentList {
		err = message_service.SendMessage(application_type.APP, studentId, senderMessage)
		if err != nil {
			return err
		}
		err = message_service.SendMessage(application_type.WEB, studentId, senderMessage)
		if err != nil {
			return err
		}
	}
	return nil
}

func RequireAccess(senderMessage dto.Message) error {
	// fill in asker's information
	userInfo, err := user_info_service.GetUserInfo(senderMessage.SenderId)
	if err != nil {
		return err
	}
	senderMessage.Message = fmt.Sprintf(`{"userAvatar": "%s", "userNickname": "%s"}`, userInfo.AvatorUrl, userInfo.Nickname)

	// send asking message
	// don't know what application type does mentor use, so send message to all application type
	// user can only use one application to open classroom afterall
	mentorId, err := classroom_repository.FindMentorIdByClassroomId(context.TODO(), senderMessage.ClassroomId)
	err = message_service.SendMessageToAllApplicationType(mentorId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}

func RejectAccess(senderMessage dto.Message) error {
	senderMessage.Cmd = ws_cmd.Reject
	err := message_service.SendMessage(senderMessage.ReceiverApplicationType, senderMessage.RecieverId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}

func AcceptAndSendClassroomToken(senderMessage dto.Message) error {
	err, classroomToken := classroom_registry_service.GetUserClassroomToken(senderMessage.SenderId)
	if err != nil {
		return err
	}

	senderMessage.Message = classroomToken
	err = message_service.SendMessage(senderMessage.ReceiverApplicationType, senderMessage.RecieverId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}

func JoinClassRoom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) error {
	err, classroomId := jwt_service.VerifyTokenAndReturnClassroomId(senderMessage.Message.(string))
	if err != nil {
		log.Logger.Infof("[join room] verify token fail %+v", err)
		return errors.New("verify classroom token fail")
	}

	joinRes, err := classroom_status_repository.JoinClassroom(classroomId, senderMessage.ApplicationType, senderMessage.SenderId)
	if joinRes != "join class success" {
		log.Logger.Info(joinRes)
		return errors.New(joinRes)
	}

	senderMessage.Message = joinRes
	senderWsNode.LocatedClassroomId = classroomId
	// inform everyone member who cares classroom status, including mentor himself (mentor may be using another device to watch classroom status).
	err = message_service.SendMessageToAllStudentAndMentor(senderMessage.ClassroomId, senderMessage)
	if err != nil {
		return errors.New("verify classroom token fail")
	}
	return nil
}

func CloseClassroom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) error {
	err := classroom_status_repository.CloseClassroom(senderMessage.ClassroomId)
	if err != nil {
		return err
	}

	err = classroom_repository.DeactiveClassSetting(senderMessage.ClassroomId)
	if err != nil {
		return err
	}

	senderWsNode.LocatedClassroomId = ""

	// inform everyone member who cares classroom status, including mentor himself (mentor may be using another device to watch classroom status).
	message_service.SendMessageToAllApplicationType(senderMessage.SenderId, senderMessage)
	studentList, err := list_service.GetStudentListByClassroomId(senderMessage.ClassroomId)
	if err != nil {
		return err
	}
	for _, studentId := range studentList {
		err = message_service.SendMessageToAllApplicationType(studentId, senderMessage)
		if err != nil {
			return err
		}
	}
	return nil
}

func LeaveClassroom(senderMessage dto.Message, senderWsNode *dto.WebSocketNode) error {
	err := message_service.SendMessageToSpecificClassroom(senderMessage.ClassroomId, senderMessage)
	if err != nil {
		return err
	}

	err = classroom_status_repository.LeaveClassroom(senderMessage.ClassroomId, senderMessage.ApplicationType, senderMessage.SenderId)
	if err != nil {
		return err
	}

	// inform everyone member who cares classroom status, including mentor himself (mentor may be using another device to watch classroom status).
	// message_service.SendMessageToAllApplicationType(senderMessage.SenderId, senderMessage)
	// studentList, err := list_service.GetStudentListByClassroomId(senderMessage.ClassroomId)
	// if err != nil {
	// 	return err
	// }
	// for _, studentId := range studentList {
	// 	err = message_service.SendMessageToAllApplicationType(studentId, senderMessage)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	senderWsNode.LocatedClassroomId = ""

	err = message_service.SendMessageToAllStudentAndMentor(senderMessage.ClassroomId, senderMessage)
	if err != nil {
		return errors.New("send leave classroom message fail!")
	}
	return nil
}
