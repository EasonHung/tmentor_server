package user_system_response

type GetUserInfoRes struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    GetUserInfoResData `json:"data"`
}

type GetUserInfoResData struct {
	AvatorUrl      string            `json:"avatorUrl"`
	UserId         string            `json:"userId"`
	Nickname       string            `json:"nickname"`
	AboutMe        string            `json:"aboutMe"`
	Education      []EducationVo     `json:"education"`
	JobExperiences []JobExperienceVo `json:"jobExperiences"`
	Gender         string            `json:"gender"`
	Profession     []string          `json:"profession"`
	Fields         []string          `json:"fields"`
	PictureUrl     string            `json:"pictureUrl"`
	MentorSkill    []string          `json:"mentorSkill"`
}

type JobExperienceVo struct {
	CompanyName string `json:"companyName"`
	JobName     string `json:"jobName"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

type EducationVo struct {
	SchoolName string `json:"schoolName"`
	Subject    string `json:"subject"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}
