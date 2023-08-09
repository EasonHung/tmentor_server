package classroom_info_service

import (
	"context"
	"mentor/classroom/constants/classroom_status_enum"
	"mentor/classroom/domain/classroom/classroom_repository"
	"mentor/classroom/domain/classroom/dto"
	"mentor/classroom/domain/classroom_status/classroom_status_repository"
	"mentor/classroom/domain/ownership/ownership_repository"
	"mentor/classroom/dto/res"
)

func GetClassroomStatus(claassroomId string) (res.GetClassroomStatusRes, error) {
	mentorId, err := classroom_repository.FindMentorIdByClassroomId(context.TODO(), claassroomId)
	if err != nil {
		return res.GetClassroomStatusRes{}, err
	}

	classroomStatus, err := classroom_status_repository.GetClassroomStatus(claassroomId, mentorId)
	if err != nil {
		return res.GetClassroomStatusRes{}, err
	}

	classSetting := classroom_dto.ClassSetting{}
	if classroomStatus == classroom_status_enum.ONLINE || classroomStatus == classroom_status_enum.IN_CLASS {
		classSetting, err = classroom_repository.GetActiveClassSetting(claassroomId)
		if err != nil {
			return res.GetClassroomStatusRes{}, err
		}
	}

	return res.NewGetClassroomStatusRes(classroomStatus, classSetting), nil
}

func GetClassroomId(userId string) (string, error) {
	classroomId, err := classroom_repository.FindClassroomIdByMentorId(context.TODO(), userId)
	if err != nil {
		return "", err
	}

	return classroomId, nil
}

func GetStudentCount(userId string) (int, error) {
	studentList, err := ownership_repository.GetStudentList(context.TODO(), userId)
	if err != nil {
		return 0, err
	}

	return len(studentList), nil
}
