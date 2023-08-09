package list_service

import (
	"context"
	"mentor/classroom/constants/classroom_status_enum"
	"mentor/classroom/domain/classroom/classroom_repository"
	classroom_dto "mentor/classroom/domain/classroom/dto"
	"mentor/classroom/domain/classroom_status/classroom_status_repository"
	"mentor/classroom/domain/ownership/ownership_repository"
	"mentor/classroom/dto/res"
	"mentor/classroom/service/user_info_service"

	"github.com/pkg/errors"
)

func GetClassroomList(userId string) ([]res.TeacherItem, error) {
	classroomList, err := ownership_repository.GetClassroomList(context.TODO(), userId)
	if err != nil {
		return nil, err
	}

	resList := make([]res.TeacherItem, 0)
	for _, classroomInfo := range classroomList {
		classroomStatus, err := classroom_status_repository.GetClassroomStatus(classroomInfo.ClassroomId, classroomInfo.MentorId)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		mentorInfo, err := user_info_service.GetUserInfo(classroomInfo.MentorId)
		if err != nil {
			return nil, err
		}

		classSetting := classroom_dto.ClassSetting{}
		if classroomStatus == classroom_status_enum.ONLINE || classroomStatus == classroom_status_enum.IN_CLASS {
			classSetting, err = classroom_repository.GetActiveClassSetting(classroomInfo.ClassroomId)
			if err != nil {
				return nil, errors.New(err.Error())
			}
		}

		resItem := res.TeacherItem{}
		resItem.FromDto(mentorInfo, classSetting, classroomStatus, classroomInfo.ClassroomId)
		resList = append(resList, resItem)
	}

	return resList, nil
}

func GetStudentListByClassroomId(classroomId string) ([]string, error) {
	userId, err := classroom_repository.FindMentorIdByClassroomId(context.TODO(), classroomId)
	if err != nil {
		return nil, err
	}

	studentList, err := ownership_repository.GetStudentList(context.TODO(), userId)
	if err != nil {
		return nil, err
	}
	return studentList, nil
}

func GetMentorIdAndStudentListByClassroomId(classroomId string) (string, []string, error) {
	mentorId, err := classroom_repository.FindMentorIdByClassroomId(context.TODO(), classroomId)
	if err != nil {
		return "", nil, err
	}

	studentList, err := ownership_repository.GetStudentList(context.TODO(), mentorId)
	if err != nil {
		return "", nil, err
	}
	return mentorId, studentList, nil
}
