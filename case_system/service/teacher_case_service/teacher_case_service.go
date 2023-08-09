package teacher_case_service

import (
	"case_system/dao/teacher_case_dao"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const PAGE_PER_SIZE int64 = 30

func AddTeacherCase(avatar string, userId string, nickname string, postTime time.Time, title string, content string, pictureUrl string) error {
	teacherCaseId := bson.NewObjectId().Hex()
	applyIds := make([]string, 0)
	entity := teacher_case_dao.TeacherCase{
		TeacherCaseId: teacherCaseId,
		ApplyIds:      applyIds,
		AvatarUrl:     avatar,
		UserId:        userId,
		Nickname:      nickname,
		PostTime:      postTime,
		Title:         title,
		Content:       content,
		PictureUrl:    pictureUrl,
	}

	_, err := teacher_case_dao.InsertOne(entity)

	return err
}

func GetTeacherCaseByPage(page int64) ([]teacher_case_dao.TeacherCase, error) {
	cases, err := teacher_case_dao.FindWithPagination(page, PAGE_PER_SIZE)
	if err != nil {
		return nil, err
	}

	return cases, nil
}
