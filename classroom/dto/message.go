package dto

import (
	"mentor/classroom/constants/ws_cmd"
	"time"
)

type Message struct {
	ClassroomId             string      `json:"classroomId"`
	Cmd                     string      `json:"cmd"`
	SenderId                string      `json:"senderId"`
	RecieverId              string      `json:"recieverId"`
	ApplicationType         string      `json:"applicationType"`
	ReceiverApplicationType string      `json:"receiverApplicationType"`
	Time                    time.Time   `json:"time"`
	Message                 interface{} `json:"message"`
}

func (this *Message) FromOpenClassroomCmdToSuccessCmd() Message {
	return Message{
		ClassroomId:             this.ClassroomId,
		Cmd:                     this.Cmd,
		SenderId:                this.SenderId,
		RecieverId:              this.RecieverId,
		ApplicationType:         this.ApplicationType,
		ReceiverApplicationType: this.ReceiverApplicationType,
		Time:                    this.Time,
		Message:                 "open class success",
	}
}

func (this *Message) ToErrorCmd(errorMessage string) {
	this.Cmd = ws_cmd.UnexpectError
	this.Message = errorMessage
}

func GenerateCloseClassroomCmd(classroomId string, userId string, applicationType string) Message {
	return Message{
		ClassroomId: classroomId,
		Cmd: ws_cmd.CloseRoom,
		SenderId: userId,
		ApplicationType: applicationType,
	}
}

func GenerateLeaveClassroomCmd(classroomId string, userId string, applicationType string) Message {
	return Message{
		ClassroomId: classroomId,
		Cmd: ws_cmd.LeaveRoom,
		SenderId: userId,
		ApplicationType: applicationType,
	}
}
