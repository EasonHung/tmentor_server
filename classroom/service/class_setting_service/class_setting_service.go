package class_setting_service

import (
	"context"
	"mentor/classroom/domain/classroom"
	"mentor/classroom/domain/classroom/dto"
	"mentor/classroom/domain/classroom/classroom_repository"
	"mentor/classroom/internal_error"

	"github.com/pkg/errors"
)

const SETTING_COUNT_LIMIT = 2

func AddClassSetting(userId string, newClassSetting classroom.ClassSetting) error {
	storedClassSettingList, err := classroom_repository.GetStoredClassSetting(context.TODO(), userId)
	if err != nil {
		return err
	}

	// check if excess setting limit count
	if len(storedClassSettingList) >= SETTING_COUNT_LIMIT {
		return errors.WithStack(internal_error.OutOfLimitError{})
	}

	// check if already have the same setting name
	for _, classSetting := range storedClassSettingList {
		if classSetting.SettingName == newClassSetting.SettingName {
			return errors.WithStack(internal_error.DuplicateKeyError{})
		}
	}

	err = classroom_repository.AddNewClassSetting(context.TODO(), userId, newClassSetting)
	if err != nil {
		return errors.Wrap(err, "error push class settings")
	}

	return nil
}

func UpdateClassSetting(userId string, updatedClassSetting classroom.ClassSetting) error {
	err := classroom_repository.UpdateClassSetting(context.TODO(), userId, updatedClassSetting)
	if err != nil  {
		return err
	}
	return nil
}

func DeleteClassSetting(userId string, classSettingName string) error {
	err := classroom_repository.DeleteClassSetting(context.TODO(), userId, classSettingName)
	if err != nil {
		return err
	}
	return nil
}

func GetClassSettingList(userId string) ([]classroom_dto.ClassSetting, error){
	storedClassSettingList, err := classroom_repository.GetStoredClassSetting(context.TODO(), userId)
	if err != nil {
		return nil, err
	}
	return storedClassSettingList, nil
}