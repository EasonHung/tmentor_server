package class_info_service

import (
	"context"
	"mentor/classroom/dao/class_info_dao"
	"mentor/classroom/dao/class_info_dao/enum/class_info_status"
	"mentor/classroom/domain/class/class_repository"
	"mentor/classroom/domain/class/constants/class_status"
	class_dto "mentor/classroom/domain/class/dto"

	"github.com/pkg/errors"
)

func GetLastUnfinishedClassInfo(classroomId string, studentId string) (error, class_dto.LastClassInfo) {
	err, res := class_repository.FindLastClass(context.TODO(), classroomId, studentId)
	if err != nil {
		switch err.Error() {
		case "mongo: no documents in result":
			return errors.New("no data"), res
		default:
			return err, res
		}
	}

	if res.Status == class_status.Finish {
		return errors.New("no data"), res
	}
	return nil, res
}

func FinishClass(classId string) (error, class_info_dao.ClassInfo) {
	classInfo, err := class_info_dao.FindByClassId(classId)
	if err != nil {
		return err, classInfo
	}

	class_info_dao.UpdateStatusByClassId(classId, class_info_status.Finish)

	return nil, classInfo
}
