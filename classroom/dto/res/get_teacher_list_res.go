package res

import (
	classroom_dto "mentor/classroom/domain/classroom/dto"
	"mentor/classroom/internal_api/user_system_api/user_system_response"
)

type TeacherItem struct {
	ClassroomId       string `json:"classroomId"`
	TeacherId         string `json:"teacherId"`
	Title             string `json:"title"`
	Status            string `json:"status"`
	TeacherAvatar     string `json:"teacherAvatar"`
	TeacherNickname   string `json:"teacherNickname"`
	TeacherProfession []string `json:"teacherProfession"`
	ClassTitle        string `json:"classTitle"`
	ClassDesc         string `json:"classDesc"`
	ClassTime         int    `json:"classTime"`
	ClassPoints       int    `json:"classPoints"`
}

func (this *TeacherItem) FromDto(userInfo user_system_response.GetUserInfoResData, classSetting classroom_dto.ClassSetting, classroomStatus string, classroomId string) {
	this.ClassroomId = classroomId
	this.TeacherId = userInfo.UserId
	this.Title = classSetting.Title
	this.Status = classroomStatus
	this.TeacherAvatar = userInfo.AvatorUrl
	this.TeacherNickname = userInfo.Nickname
	this.TeacherProfession = userInfo.Profession
	this.ClassTitle = classSetting.Title
	this.ClassDesc = classSetting.Desc
	this.ClassTime = classSetting.ClassTime
	this.ClassPoints = classSetting.ClassPoints
}
