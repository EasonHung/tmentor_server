package class_process_service

import (
	"context"
	"encoding/json"
	"mentor/classroom/constants/ws_cmd"
	"mentor/classroom/domain/class/class_repository"
	"mentor/classroom/domain/class/constants/class_status"
	"mentor/classroom/dto"
	"mentor/classroom/internal_api/finance_system_api"
	"mentor/classroom/internal_api/finance_system_api/finance_system_error_code"
	"mentor/classroom/service/message_service"

	"github.com/pkg/errors"
)

func AskForAcceptance(classId string, senderMessage dto.Message) error {
	err, classInfo := class_repository.FindByClassId(context.TODO(), classId)
	if err != nil {
		return err
	}

	// insert class info json string to sender's message
	classInfoJsonStr, err := json.Marshal(classInfo)
	if err != nil {
		return errors.WithStack(err)
	}
	senderMessage.Message = string(classInfoJsonStr)

	err = message_service.SendMessageToAllApplicationType(senderMessage.RecieverId, senderMessage)
	if err != nil {
		return err
	}

	return nil
}

func InitClass(classroomId string, mentorId string, studentId string, classTitle string, classDesc string, points int, classTime int) (error, string) {
	err, classId := class_repository.InitClass(context.TODO(), classroomId, mentorId, studentId, classTitle, classDesc, points, classTime)
	if err != nil {
		return err, ""
	}
	return nil, classId
}

func ReimburseClass(classId string) error {
	err, classInfo := class_repository.FindByClassId(context.TODO(), classId)
	if err != nil {
		return err
	}

	err = class_repository.ReimburseClass(context.TODO(), classId, classInfo.RemainTime)
	if err != nil {
		return err
	}
	return nil
}

func StartClass(classId string) error {
	// check class status
	err, classInfo := class_repository.FindByClassId(context.TODO(), classId)
	if err != nil {
		return err
	}
	if classInfo.Status != class_status.Init && classInfo.Status != class_status.Reimburse {
		return errors.New("class has started")
	}

	err = class_repository.StartClass(context.TODO(), classId)
	if err != nil {
		return err
	}
	return nil
}

func AcceptClass(classId string, senderMessage dto.Message) error {
	err, classInfo := class_repository.FindByClassId(context.TODO(), classId)
	if err != nil {
		return err
	}

	// todo: 可以多幾個付款狀況的 class status
	err, response := finance_system_api.PayClassBill(classInfo.ClassId, classInfo.Points, classInfo.MentorId, classInfo.StudentId)
	if err != nil {
		return err
	}
	if response.Code != finance_system_error_code.SUCCESS {
		return errors.New(response.Message)
	}

	err = StartClass(classId)
	if err != nil {
		return err
	}

	senderMessage.Message = classInfo.ClassId
	senderMessage.Cmd = ws_cmd.StartClass
	err = message_service.SendMessageToSpecificClassroom(senderMessage.ClassroomId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}

func FinishClass(classId string, classroomId string, senderMessage dto.Message) error {
	err := class_repository.FinishClass(context.TODO(), classId)
	if err != nil {
		return err
	}

	err = message_service.SendMessageToSpecificClassroom(classroomId, senderMessage)
	if err != nil {
		return err
	}

	return nil
}
