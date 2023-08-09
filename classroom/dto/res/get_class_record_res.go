package res

import (
	class_dto "mentor/classroom/domain/class/dto"
	"time"
)

type GetClassRecordRes struct {
	ClassRecordList []GetClassRecordResItem `json:"classRecordList"`
}

type GetClassRecordResItem struct {
	ClassId     string                 `json:"classId"`
	ClassroomId string                 `json:"classroomId"`
	MentorInfo  ClassRecordMentorInfo  `json:"mentorInfo"`
	StudentInfo ClassRecordStudentInfo `json:"studentInfo"`
	Status      string                 `json:"status"`
	Points      int                    `json:"points"`
	Title       string                 `json:"title"`
	Desc        string                 `json:"desc"`
	ClassTime   int                    `json:"classTime"`
	RemainTime  int                    `json:"remainTime"`
	StartTime   time.Time              `json:"startTime"`
}

type ClassRecordMentorInfo struct {
	MentorId         string `json:"mentorId"`
	MentorNickname   string `json:"mentorNickname"`
	MentorProfession string `json:"mentorProfession"`
}

type ClassRecordStudentInfo struct {
	StudentId         string `json:"studentId"`
	StudentNickname   string `json:"sutdentNickname"`
	StudentProfession string `json:"studentProfession"`
}

func ConvertFromClassRecordDto(classRecordDtoList []class_dto.ClassRecord) GetClassRecordRes {
	res := make([]GetClassRecordResItem, 0)

	for _, classRecord := range classRecordDtoList {
		getClassRecordResItem := GetClassRecordResItem{
			ClassId:     classRecord.ClassId,
			ClassroomId: classRecord.ClassroomId,
			MentorInfo: ClassRecordMentorInfo{
				MentorId: classRecord.MentorId,
			},
			StudentInfo: ClassRecordStudentInfo{
				StudentId: classRecord.StudentId,
			},
			Status:     classRecord.Status,
			Points:     classRecord.Points,
			Title:      classRecord.ClassTitle,
			Desc:       classRecord.ClassDesc,
			ClassTime:  classRecord.ClassTime,
			StartTime:  classRecord.StartTime,
			RemainTime: classRecord.RemainTime,
		}
		res = append(res, getClassRecordResItem)
	}

	return GetClassRecordRes{
		ClassRecordList: res,
	}
}
