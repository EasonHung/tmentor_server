package classroom

import (
	"mentor/classroom/dto/req"

	"gopkg.in/mgo.v2/bson"
)

type Classroom struct {
	Id               string         `bson:"_id,omitempty"`
	ClassroomId      string         `bson:"classroomId"`
	MentorId         string         `bson:"mentorId"`
	ClassSettingList []ClassSetting `bson:"classSettingList"`
}

type ClassSetting struct {
	SettingName string `bson:"settingName"`
	Title       string `bson:"title"`
	Desc        string `bson:"desc"`
	ClassTime   int    `bson:"classTime"`
	ClassPoints int    `bson:"classPoints"`
}

func NewClassroom(mentorId string) Classroom {
	classSetting := make([]ClassSetting, 0)
	classroomId := bson.NewObjectId().Hex()

	return Classroom{
		ClassroomId: classroomId,
		MentorId: mentorId,
		ClassSettingList: classSetting,
	}
}

func NewClassSettingFromAddClassSettingReq(addClassSettingReq req.AddClassSettingReq) ClassSetting {
	return ClassSetting {
		SettingName: addClassSettingReq.SettingName,
		Title: addClassSettingReq.Title,
		Desc: addClassSettingReq.Desc,
		ClassTime: addClassSettingReq.ClassTime,
		ClassPoints: addClassSettingReq.ClassPoints,
	}
}

func NewClassSettingFromUpdateClassSettingReq(updateClassSettingReq req.UpdateClassSettingReq) ClassSetting {
	return ClassSetting {
		SettingName: updateClassSettingReq.SettingName,
		Title: updateClassSettingReq.Title,
		Desc: updateClassSettingReq.Desc,
		ClassTime: updateClassSettingReq.ClassTime,
		ClassPoints: updateClassSettingReq.ClassPoints,
	}
}