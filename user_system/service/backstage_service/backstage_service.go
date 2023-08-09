package backstage_service

import (
	"context"
	"mentor_app/user_system/db/connection"
	"mentor_app/user_system/db/dao/profession_catelog_dao"
	"mentor_app/user_system/db/dao/user_info_dao"
	"mentor_app/user_system/mentor_redis"
	"mentor_app/user_system/service/service_constant"

	"github.com/pkg/errors"
)

func GetTokenBlackList() (error, []string) {
	redisClient := mentor_redis.Client

	list, err := redisClient.SMembers(service_constant.USER_BLACKLIST).Result()
	if err != nil {
		return errors.Wrap(err, "error get black list"), nil
	}

	return nil, list
}

func DeleteUserIdIntoBlackList(userId string) error {
	redisClient := mentor_redis.Client

	// remove all user id in black list
	err := redisClient.SRem(service_constant.USER_BLACKLIST, userId).Err()
	if err != nil {
		return errors.Wrap(err, "error delete user black list set")
	}

	return nil
}

func AddUserIdIntoBlackList(userId string) error {
	redisClient := mentor_redis.Client

	err := redisClient.SAdd(service_constant.USER_BLACKLIST, userId).Err()
	if err != nil {
		return errors.Wrap(err, "error push user black list set")
	}

	return nil
}

func GetProfessionCategoryList() ([]string, error) {
	categoryList, err := profession_catelog_dao.DistinctCategory()
	if err != nil {
		err = errors.Wrap(err, "error distinct categories")
		return nil, err
	}

	return categoryList, nil
}

func GetUnCategorizeNames() ([]string, error) {
	unCategorizeNames := make([]string, 0)
	catelogInfos, err := profession_catelog_dao.FindByCategory("")
	if err != nil {
		err = errors.Wrap(err, "error get uncategorize infos")
		return nil, err
	}

	for _, info := range catelogInfos {
		unCategorizeNames = append(unCategorizeNames, info.Name)
	}

	return unCategorizeNames, nil
}

func AddProfessionCategory(name string, category string) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		newCategory := profession_catelog_dao.ProfessionCategory{
			Name:     name,
			Category: category,
		}
		_, err := profession_catelog_dao.CreateOneWithTx(ctx, newCategory)
		if err != nil {
			err = errors.Wrap(err, "error new category")
			return nil, err
		}

		err = user_info_dao.PushProfessionCategoryWithTx(ctx, name, category)
		if err != nil {
			err = errors.Wrap(err, "error push profession category")
			return nil, err
		}

		return nil, nil
	}

	_, err := connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	if err != nil {
		err = errors.Wrap(err, "error change category")
		return err
	}
	return nil
}

func UpdateProfessionCatelog(name string, category string, oldCategory string) error {
	categoryInfos, err := profession_catelog_dao.FindByName(name)
	if err != nil {
		errors.Wrap(err, "error get catelog info")
		return err
	}

	for _, categoryInfo := range categoryInfos {
		if categoryInfo.Category == "" {
			firstCategorize(name, category)
			continue
		}
		if categoryInfo.Category == oldCategory {
			modifyCategory(name, category, oldCategory)
		}
	}

	return nil
}

func modifyCategory(name string, newCategory string, oldCategory string) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		err := user_info_dao.PullProfessionCategoryWithTx(ctx, name, oldCategory)
		if err != nil {
			err = errors.Wrap(err, "error pull profession category")
			return nil, err
		}

		err = user_info_dao.PushProfessionCategoryWithTx(ctx, name, newCategory)
		if err != nil {
			err = errors.Wrap(err, "error push profession category")
			return nil, err
		}

		err = profession_catelog_dao.UpdateCategoryWithTx(ctx, name, newCategory, oldCategory)
		if err != nil {
			err = errors.Wrap(err, "error change category")
			return nil, err
		}

		return nil, nil
	}

	_, err := connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	return err
}

func firstCategorize(name string, category string) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		err := user_info_dao.PushProfessionCategoryWithTx(ctx, name, category)
		if err != nil {
			err = errors.Wrap(err, "error push profession category")
			return nil, err
		}

		err = profession_catelog_dao.UpdateCategoryWithTx(ctx, name, category, "")
		if err != nil {
			err = errors.Wrap(err, "error change category")
			return nil, err
		}

		return nil, nil
	}

	_, err := connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)
	return err
}
