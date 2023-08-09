package vo

import (
	"case_system/dao/teacher_case_dao"
	"time"
)

type GetTeacherCaseResVo struct {
	Data []GetTeacherCaseResVoItem `json:"data"`
}

type GetTeacherCaseResVoItem struct {
	TeacherCaseId string    `json:"teacherCaseId"`
	ApplyIds      []string  `json:"applyIds"`
	AvatarUrl     string    `json:"avatarUrl"`
	UserId        string    `json:"userId"`
	Nickname      string    `json:"nickname"`
	PostTime      time.Time `json:"postTime"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	PictureUrl    string    `json:"pictureUrl"`
}

func (this *GetTeacherCaseResVo) TeacherCaseListConvertor(teacherCaseList []teacher_case_dao.TeacherCase) {
	teacherCaseListRes := make([]GetTeacherCaseResVoItem, 0)

	for _, teacherCaseEntity := range teacherCaseList {
		resItem := GetTeacherCaseResVoItem{
			TeacherCaseId: teacherCaseEntity.TeacherCaseId,
			ApplyIds:      teacherCaseEntity.ApplyIds,
			AvatarUrl:     teacherCaseEntity.AvatarUrl,
			UserId:        teacherCaseEntity.UserId,
			Nickname:      teacherCaseEntity.Nickname,
			PostTime:      teacherCaseEntity.PostTime,
			Title:         teacherCaseEntity.Title,
			Content:       teacherCaseEntity.Content,
			PictureUrl:    teacherCaseEntity.PictureUrl,
		}
		teacherCaseListRes = append(teacherCaseListRes, resItem)
	}

	this.Data = teacherCaseListRes
}
