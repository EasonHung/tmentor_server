package class_record_service

import (
	"context"
	"mentor/classroom/domain/class/class_repository"
	"mentor/classroom/dto/res"
	"mentor/classroom/service/user_info_service"
)

func GetClassRecord(userId string) (error, res.GetClassRecordRes) {
	err, classRecordList := class_repository.FindClassRecord(context.TODO(), userId)
	if err != nil {
		return err, res.GetClassRecordRes{}
	}

	getClassRecordRes := res.ConvertFromClassRecordDto(classRecordList)

	for index, classRecord := range getClassRecordRes.ClassRecordList {
		mentorInfo, err := user_info_service.GetUserInfo(classRecord.MentorInfo.MentorId)
		if err != nil {
			return err, getClassRecordRes
		}
		studentInfo, err := user_info_service.GetUserInfo(classRecord.StudentInfo.StudentId)
		if err != nil {
			return err, getClassRecordRes
		}

		getClassRecordRes.ClassRecordList[index].MentorInfo.MentorNickname = mentorInfo.Nickname
		getClassRecordRes.ClassRecordList[index].MentorInfo.MentorProfession = mentorInfo.Profession[0]
		getClassRecordRes.ClassRecordList[index].StudentInfo.StudentNickname = studentInfo.Nickname
		getClassRecordRes.ClassRecordList[index].StudentInfo.StudentProfession = studentInfo.Profession[0]
	}

	return nil, getClassRecordRes
}