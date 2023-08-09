package classroom_info_dao

import (
	"context"
	"mentor/classroom/dao/db_connection"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var collection *qmgo.Collection

func init() {
	collection = db_connection.MONGO_CLIENT.Database("classroom").Collection("classroom_info")
}

func FindByClassroomId(classroomId string) (ClassroomInfo, error) {
	result := ClassroomInfo{}
	err := collection.Find(context.Background(), bson.M{"classroomId": classroomId}).One(&result)
	return result, err
}
