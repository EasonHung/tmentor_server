package classroom_registry_service

import (
	"context"
	"mentor/classroom/db_connection"
	"mentor/classroom/domain/classroom/classroom_repository"
	"mentor/classroom/domain/ownership/ownership_repository"
	"mentor/classroom/service/jwt_service"

	"github.com/pkg/errors"
)

func GetUserClassroomToken(mentorId string) (error, string) {
	classroomId, err := classroom_repository.FindClassroomIdByMentorId(context.TODO(), mentorId)
	if err != nil {
		return err, ""
	}

	token, err := jwt_service.GenerateClasstoken(mentorId, classroomId)
	if err != nil {
		return err, ""
	}

	return nil, token
}

func EnrollClassroom(classroomToken string, userId string) error {
	err, classroomId := jwt_service.VerifyTokenAndReturnClassroomId(classroomToken)
	if err != nil {
		return err
	}

	mentorId, err := classroom_repository.FindMentorIdByClassroomId(context.TODO(), classroomId)

	transaction := func(ctx context.Context) (interface{}, error) {
		err := ownership_repository.EnrollNewClassroom(ctx, userId, classroomId, mentorId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err = db_connection.MONGO_CLIENT.DoTransaction(context.TODO(), transaction)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
