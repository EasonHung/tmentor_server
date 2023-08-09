package web_rtc_service

import (
	"mentor/classroom/dto"
	"mentor/classroom/service/message_service"
)

func SendOffer(senderMessage dto.Message) error {
	err := message_service.SendMessageToOtherMembersInClassroom(senderMessage.ClassroomId, senderMessage.SenderId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}

func SendAnswer(senderMessage dto.Message) error {
	err := message_service.SendMessageToOtherMembersInClassroom(senderMessage.ClassroomId, senderMessage.SenderId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}

func SendCandidate(senderMessage dto.Message) error {
	err := message_service.SendMessageToOtherMembersInClassroom(senderMessage.ClassroomId, senderMessage.SenderId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}
