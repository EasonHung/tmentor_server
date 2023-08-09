package student_case_dao

import "time"

type StudentCase struct {
	Id            string    `bson:"_id,omitempty"`
	StudentCaseId string    `bson:"studentCaseId"`
	BidInfoIds    []string  `bson:"bidInfoIds"`
	AvatarUrl     string    `bson:"avatarUrl"`
	UserId        string    `bson:"userId"`
	Nickname      string    `bson:"nickname"`
	PostTime      time.Time `bson:"postTime"`
	Title         string    `bson:"title"`
	Content       string    `bson:"content"`
	PictureUrl    []string  `bson:"pictureUrl"`
}
