package class_time_service

import (
	"context"
	"mentor/classroom/domain/class/class_repository"
	"mentor/classroom/domain/class/constants/class_status"
	"time"

	"github.com/pkg/errors"
)

func GetStartTime(classId string) (error, time.Time) {
	err, classInfo := class_repository.FindByClassId(context.TODO(), classId)
	if err != nil {
		return err, time.Now()
	}

	if classInfo.Status != class_status.Start {
		return errors.New("wrong status"), time.Now()
	}
	
	return nil, classInfo.StartTime
}

func ClockOn(classId string, userId string) error {
	err, classInfo := class_repository.FindByClassId(context.TODO(), classId)
	if err != nil {
		return err
	}

	timeGoesBy := time.Now().Sub(classInfo.StartTime).Minutes()
	remainTime := classInfo.ClassTime - int(timeGoesBy)

	err = class_repository.ClockOn(context.TODO(), classId, userId, remainTime)
	if err != nil {
		return err
	}
	return nil
}