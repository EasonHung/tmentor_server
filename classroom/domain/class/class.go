package class

import (
	"mentor/classroom/domain/class/constants/class_status"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Class struct {
	Id          string        `bson:"_id,omitempty"`
	ClassId     string        `bson:"classId"`
	ClassroomId string        `bson:"classroomId"`
	MentorId    string        `bson:"mentorId"`
	StudentId   string        `bson:"studentId"`
	Status      string        `bson:"status"`
	Title       string        `bson:"title"`
	Desc        string        `bson:"desc"`
	Points      int           `bson:"points"`
	ClassTime   int           `bson:"classTime"`
	RemainTime  int           `bson:"remainTime"`
	StartTime   time.Time     `bson:"startTime"`
	ClockRecord []ClockRecord `bson:"clockRecord"`
}

type ClockRecord struct {
	UserId    string    `bson:"userId"`
	ClockTime time.Time `bson:"clockTime"`
}

func NewClass(classroomId string, mentorId string, studentId string, classTitle string, classDesc string, points int, classTime int) Class {
	classId := bson.NewObjectId().Hex()
	emptyClockRecord := make([]ClockRecord, 0)
	return Class{
		ClassId:     classId,
		ClassroomId: classroomId,
		MentorId:    mentorId,
		StudentId:   studentId,
		Status:      class_status.Init,
		Title:       classTitle,
		Desc:        classDesc,
		Points:      points,
		ClassTime:   classTime,
		RemainTime:  classTime,
		ClockRecord: emptyClockRecord,
	}
}
