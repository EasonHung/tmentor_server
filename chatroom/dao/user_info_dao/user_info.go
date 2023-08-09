package user_info_dao

import "time"

type UserInfo struct {
	AvatorUrl   string    `bson:"avatorUrl"`
	UserId      string    `bson:"userId"`
	FcmToken    string    `bson:"fcmToken"`
	Nickname    string    `bson:"nickname"`
	AboutMe     string    `bson:"aboutMe"`
	Education   string    `bson:"education"`
	Gender      string    `bson:"gender"`
	Profession  string    `bson:"profession"`
	PictureUrl  string    `bson:"pictureUrl"`
	MentorSkill []string  `bson:"mentorSkill"`
	CreateTime  time.Time `bson:"createTime"`
	ModifyTime  time.Time `bson:"modifyTime"`
}
