package vo

import "mentor_app/user_system/db/dao/user_info_dao"

type GetUserInfoRes struct {
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

func (this *GetUserInfoRes) UserInfoDaoConvertor(dao user_info_dao.UserInfo) {
	experienceVoList := make([]JobExperienceVo, 0)
	for _, jobExperience := range dao.JobExperiences {
		experience := JobExperienceVo{
			CompanyName: jobExperience.CompanyName,
			JobName:     jobExperience.JobName,
			StartTime:   jobExperience.StartTime,
			EndTime:     jobExperience.EndTime,
		}
		experienceVoList = append(experienceVoList, experience)
	}

	educationVoList := make([]EducationVo, 0)
	for _, educationDao := range dao.Education {
		educationVo := EducationVo{
			SchoolName: educationDao.SchoolName,
			Subject:    educationDao.Subject,
			StartTime:  educationDao.StartTime,
			EndTime:    educationDao.EndTime,
		}
		educationVoList = append(educationVoList, educationVo)
	}

	this.AvatorUrl = dao.AvatorUrl
	this.UserId = dao.UserId
	this.Nickname = dao.Nickname
	this.AboutMe = dao.AboutMe
	this.Education = educationVoList
	this.Gender = dao.Gender
	this.Profession = dao.Profession
	this.Fields = dao.Fields
	this.PictureUrl = dao.PictureUrl
	this.MentorSkill = dao.MentorSkill
	this.JobExperiences = experienceVoList
}
