package evaluation_posts_dao

import (
	"context"
	"evaluation_system/dao/db_connection"
	"time"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

const PAGE_SIZE = 20

func init() {
	collection = db_connection.MONGO_CLIENT.Database("evaluation_system").Collection("evaluation_posts")
}

func InsertOneWithTx(ctx context.Context, obj EvaluationPost) (*qmgo.InsertOneResult, error) {
	obj.CreateTime = time.Now()
	obj.ModifyTime = time.Now()
	result, err := collection.InsertOne(ctx, obj)
	return result, err
}

func FindByToUserWithPage(page int64, toUserId string) ([]EvaluationPost, error) {
	batch := []EvaluationPost{}
	offset := page * PAGE_SIZE

	err := collection.Find(context.Background(), bson.M{"toUserId": toUserId}).Sort("-createTime").Skip(offset).Limit(page).All(&batch)
	if err != nil {
		return nil, err
	}

	return batch, nil
}
