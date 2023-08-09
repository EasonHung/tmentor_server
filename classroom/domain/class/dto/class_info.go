package class_dto

import "time"

type ClassInfo struct {
	ClassId     string    `bson:"classId" json:"classId"`
	ClassroomId string    `bson:"classroomId" json:"classroomId"`
	MentorId    string    `bson:"mentorId" json:"mentorId"`
	StudentId   string    `bson:"studentId" json:"studentId"`
	Status      string    `bson:"status" json:"status"`
	Points      int       `bson:"points" json:"points"`
	ClassTime   int       `bson:"classTime" json:"classTime"`
	RemainTime  int       `bson:"remainTime" json:"remainTime"`
	StartTime   time.Time `bson:"startTime" json:"startTime"`
}