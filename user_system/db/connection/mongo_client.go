package connection

import (
	"context"
	"mentor_app/user_system/initialize"

	"github.com/qiniu/qmgo"
)

var (
	MONGO_CLIENT *qmgo.Client
	err          error
)

func init() {
	ctx := context.Background()
	MONGO_CLIENT, err = qmgo.NewClient(ctx, &qmgo.Config{
		Uri: initialize.GLOBAL_CONFIG.Db.MongoDb.Srv})
	if err != nil {
		panic(err)
	}
}
