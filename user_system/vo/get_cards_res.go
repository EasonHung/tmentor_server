package vo

import "time"

type GetCardsRes struct {
	AvatorUrl      string            `json:"avatorUrl"`
	UserId         string            `json:"userId"`
	Nickname       string            `json:"nickname"`
	AboutMe        string            `json:"aboutMe"`
	Education      []EducationVo     `json:"education"`
	Gender         string            `json:"gender"`
	Profession     []string          `json:"profession"`
	Fields         []string          `json:"fields"`
	PictureUrl     string            `json:"pictureUrl"`
	MentorSkill    []string          `json:"mentorSkill"`
	JobExperiences []JobExperienceVo `json:"jobExperiences"`
	StudentCount   int               `json:"studentCount"`
	CreateTime     time.Time         `json:"createTime"`
	ModifyTime     time.Time         `json:"modifyTime"`
}
