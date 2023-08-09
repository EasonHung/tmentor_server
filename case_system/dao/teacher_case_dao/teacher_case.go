package teacher_case_dao

import "time"

type TeacherCase struct {
	Id            string    `bson:"_id,omitempty"`
	TeacherCaseId string    `bson:"teacherCaseId"`
	ApplyIds      []string  `bson:"applyIds"`
	AvatarUrl     string    `bson:"avatarUrl"`
	UserId        string    `bson:"userId"`
	Nickname      string    `bson:"nickname"`
	PostTime      time.Time `bson:"postTime"`
	Title         string    `bson:"title"`
	Content       string    `bson:"content"`
	PictureUrl    string    `bson:"pictureUrl"`
}
