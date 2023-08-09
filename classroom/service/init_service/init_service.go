package init_service

import (
	"context"
	"mentor/classroom/db_connection"
	"mentor/classroom/domain/classroom/classroom_repository"
	"mentor/classroom/domain/ownership/ownership_repository"
)

func InitUser(userId string) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		err := classroom_repository.CreateNewClassroom(ctx, userId)
		if err != nil {
			return nil, err
		}

		err = ownership_repository.CreateNewOwnership(ctx, userId)
		if err != nil {
			return nil, err
		}
		
		return nil, nil
	}

	_, err := db_connection.MONGO_CLIENT.DoTransaction(context.TODO(), transaction)

	return err
}