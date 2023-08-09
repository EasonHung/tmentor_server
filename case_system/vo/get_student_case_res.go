package vo

import (
	"case_system/dao/student_case_dao"
	"time"
)

type GetStudentCaseResVo struct {
	Data []GetStudentCaseResVoItem `json:"data"`
}

type GetStudentCaseResVoItem struct {
	StudentCaseId string    `json:"studentCaseId"`
	BidInfoIds    []string  `json:"bidInfoIds"`
	AvatarUrl     string    `json:"avatarUrl"`
	UserId        string    `json:"userId"`
	Nickname      string    `json:"nickname"`
	PostTime      time.Time `json:"postTime"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	PictureUrl    []string  `json:"pictureUrl"`
}

func (this *GetStudentCaseResVo) StudentCaseListConvertor(studentCaseList []student_case_dao.StudentCase) {
	studentCaseListRes := make([]GetStudentCaseResVoItem, 0)

	for _, studentCaseEntity := range studentCaseList {
		resItem := GetStudentCaseResVoItem{
			StudentCaseId: studentCaseEntity.StudentCaseId,
			BidInfoIds:    studentCaseEntity.BidInfoIds,
			AvatarUrl:     studentCaseEntity.AvatarUrl,
			UserId:        studentCaseEntity.UserId,
			Nickname:      studentCaseEntity.Nickname,
			PostTime:      studentCaseEntity.PostTime,
			Title:         studentCaseEntity.Title,
			Content:       studentCaseEntity.Content,
			PictureUrl:    studentCaseEntity.PictureUrl,
		}
		studentCaseListRes = append(studentCaseListRes, resItem)
	}

	this.Data = studentCaseListRes
}

func (this *GetStudentCaseResVoItem) StudentCaseConvertor(studentCaseEntity student_case_dao.StudentCase) {
	this.StudentCaseId = studentCaseEntity.StudentCaseId
	this.BidInfoIds = studentCaseEntity.BidInfoIds
	this.AvatarUrl = studentCaseEntity.AvatarUrl
	this.UserId = studentCaseEntity.UserId
	this.Nickname = studentCaseEntity.Nickname
	this.PostTime = studentCaseEntity.PostTime
	this.Title = studentCaseEntity.Title
	this.Content = studentCaseEntity.Content
	this.PictureUrl = studentCaseEntity.PictureUrl
}
