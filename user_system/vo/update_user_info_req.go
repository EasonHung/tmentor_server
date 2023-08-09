package vo

import "mentor_app/user_system/db/dao/user_info_dao"

type UpdateUserInfoReq struct {
	Nickname       string            `json:"nickname"`
	AboutMe        string            `json:"aboutMe"`
	Education      []EducationVo     `json:"education"`
	Gender         string            `json:"gender"`
	Profession     []string          `json:"profession"`
	JobExperiences []JobExperienceVo `json:"jobExperiences"`
	Fields         []string          `json:"fields"`
	MentorSkill    []string          `json:"mentorSkill"`
	UserStatus     string            `json:"userStatus"`
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

func (this *UpdateUserInfoReq) ValidationAndRebuildRequest() {
	if this.JobExperiences == nil {
		this.JobExperiences = make([]JobExperienceVo, 0)
	}
	if this.Education == nil {
		this.Education = make([]EducationVo, 0)
	}
	if this.Profession == nil {
		this.Profession = make([]string, 0)
	}
	if this.Fields == nil {
		this.Fields = make([]string, 0)
	}
	if this.MentorSkill == nil {
		this.MentorSkill = make([]string, 0)
	}
}

func (this *UpdateUserInfoReq) ToUserInfoDao() user_info_dao.UserInfo {
	jobExperienceList := make([]user_info_dao.JobExperience, 0)
	educationList := make([]user_info_dao.Education, 0)

	for _, jobExperience := range this.JobExperiences {
		jobDao := user_info_dao.JobExperience{
			CompanyName: jobExperience.CompanyName,
			JobName:     jobExperience.JobName,
			StartTime:   jobExperience.StartTime,
			EndTime:     jobExperience.EndTime,
		}
		jobExperienceList = append(jobExperienceList, jobDao)
	}

	for _, education := range this.Education {
		educationDao := user_info_dao.Education{
			SchoolName: education.SchoolName,
			Subject:    education.Subject,
			StartTime:  education.StartTime,
			EndTime:    education.EndTime,
		}
		educationList = append(educationList, educationDao)
	}

	return user_info_dao.UserInfo{
		Nickname:       this.Nickname,
		UserStatus:     this.UserStatus,
		AboutMe:        this.AboutMe,
		Education:      educationList,
		Gender:         this.Gender,
		JobExperiences: jobExperienceList,
		Profession:     this.Profession,
		Fields:         this.Fields,
		MentorSkill:    this.MentorSkill,
	}
}
