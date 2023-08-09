package class_info_dao

import "time"

type ClassInfo struct {
	Id            string    `bson:"_id,omitempty"`
	ClassroomId   string    `bson:"classroomId"`
	TeacherUserId string    `bson:"teacherUserId"`
	StudentUserId string    `bson:"studentUserId"`
	Status        string    `bson:"status"`
	Points        int       `bson:"points"`
	ClassTime     int       `bson:"classTime"`
	RemainTime    int       `bson:"remainTime"`
	StartTime     time.Time `bson:"startTime"`
	CreateTime    time.Time `bson:"createTime"`
	ModifyTime    time.Time `bson:"modifyTime"`
}
