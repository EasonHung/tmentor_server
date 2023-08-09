package profession_catelog_dao

import (
	"context"
	"database/sql"
	"mentor_app/user_system/db/connection"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

var collection *qmgo.Collection

func init() {
	collection = connection.MONGO_CLIENT.Database("user_info_system").Collection("profession_catelog")
}

func FindByName(name string) ([]ProfessionCategory, error) {
	batch := []ProfessionCategory{}

	err := collection.Find(context.Background(), bson.M{"name": name}).All(&batch)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return batch, nil
}

func FindByCategory(cateogry string) ([]ProfessionCategory, error) {
	batch := []ProfessionCategory{}

	err := collection.Find(context.Background(), bson.M{"category": cateogry}).All(&batch)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return batch, nil
}

func UpdateCategoryWithTx(ctx context.Context, name string, newCategory string, oldCategory string) error {
	err := collection.UpdateOne(context.Background(), bson.M{"name": name, "category": oldCategory},
		bson.M{
			"$set": bson.M{
				"category":   newCategory,
				"modifyTime": time.Now(),
			},
		})
	if err != nil {
		err = errors.Wrap(err, "error update category")
		return err
	}

	return nil
}

func CreateOne(obj ProfessionCategory) (*qmgo.InsertOneResult, error) {
	obj.CreateTime = time.Now()
	obj.ModifyTime = time.Now()
	result, err := collection.InsertOne(context.Background(), obj)
	return result, err
}

func CreateOneWithTx(ctx context.Context, obj ProfessionCategory) (*qmgo.InsertOneResult, error) {
	obj.CreateTime = time.Now()
	obj.ModifyTime = time.Now()
	result, err := collection.InsertOne(ctx, obj)
	return result, err
}

func DistinctCategory() ([]string, error) {
	batch := make([]string, 0)

	err := collection.Find(context.Background(), bson.M{}).Distinct("category", &batch)
	if err != nil {
		err = errors.Wrap(err, "error update category")
		return batch, err
	}

	return batch, nil
}
