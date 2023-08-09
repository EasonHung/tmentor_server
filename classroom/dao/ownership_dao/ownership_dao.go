package ownership_dao

import (
	"mentor/classroom/dao/db_connection"

	"github.com/qiniu/qmgo"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("classroom").Collection("ownership")
}
