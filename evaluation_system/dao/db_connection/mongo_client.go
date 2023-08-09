package db_connection

import (
	"context"

	"evaluation_system/config"

	"github.com/qiniu/qmgo"
)

var (
	MONGO_CLIENT *qmgo.Client
	err          error
)

func init() {
	ctx := context.Background()
	MONGO_CLIENT, err = qmgo.NewClient(ctx, &qmgo.Config{
		Uri: config.GLOBAL_CONFIG.Db.MongoDb.Srv})
	if err != nil {
		panic(err)
	}
}
