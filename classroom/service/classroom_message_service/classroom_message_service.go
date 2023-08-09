package classroom_message_service

import (
	"mentor/classroom/dto"
	"mentor/classroom/service/message_service"
)

func SendInstantMessage(senderMessage dto.Message) error {
	err := message_service.SendMessageToSpecificClassroom(senderMessage.ClassroomId, senderMessage)
	if err != nil {
		return err
	}
	return nil
}