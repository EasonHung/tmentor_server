package user_info_dao

import (
	"time"
)

type UserInfo struct {
	AvatorUrl      string          `bson:"avatorUrl"`
	UserId         string          `bson:"userId"`
	UserStatus     string          `bson:"userStatus"`
	LoginInfo      LoginInfo       `bson:"loginInfo"`
	ThirdParty     ThirdParty      `bson:"thirdParty"`
	FcmToken       []string        `bson:"fcmToken"`
	Nickname       string          `bson:"nickname"`
	AboutMe        string          `bson:"aboutMe"`
	Education      []Education     `bson:"education"`
	Gender         string          `bson:"gender"`
	Profession     []string        `bson:"profession"`
	JobExperiences []JobExperience `bson:"jobExperiences"`
	Fields         []string        `bson:"fields"`
	PictureUrl     string          `bson:"pictureUrl"`
	MentorSkill    []string        `bson:"mentorSkill"`
	CreateTime     time.Time       `bson:"createTime"`
	ModifyTime     time.Time       `bson:"modifyTime"`
}

type LoginInfo struct {
	Account  string `bson:"account"`
	Password string `bson:"password"`
}

type ThirdParty struct {
	ThirdPartyId   string `bson:"thirdPartyId"`
	ThirdPartyInfo string `bson:"thirdPartyInfo"`
}

type JobExperience struct {
	CompanyName string `bson:"companyName"`
	JobName     string `bson:"jobName"`
	StartTime   string `bson:"startTime"`
	EndTime     string `bson:"endTime"`
}

type Education struct {
	SchoolName string `bson:"schoolName"`
	Subject    string `bson:"subject"`
	StartTime  string `bson:"startTime"`
	EndTime    string `bson:"endTime"`
}
