package bid_info_dao

import (
	"case_system/dao/db_connection"
	"context"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("case_system").Collection("bid_info")
}

func InsertOneWithTx(ctx context.Context, obj BidInfo) (*qmgo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, obj)
	return result, err
}

func FindByBidInfoId(bidInfoId string) (BidInfo, error) {
	res := BidInfo{}

	err := collection.Find(context.Background(), bson.M{"bidInfoId": bidInfoId}).One(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func FindByBidderId(bidderId string) ([]BidInfo, error) {
	result := []BidInfo{}

	err := collection.Find(context.Background(), bson.M{"bidderId": bidderId}).Sort("-bidInfoId").All(&result)
	if err != nil {
		return nil, errors.Wrap(err, "err when get bidInfo from db")
	}
	return result, nil
}
