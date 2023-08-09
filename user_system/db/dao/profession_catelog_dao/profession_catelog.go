package profession_catelog_dao

import "time"

type ProfessionCategory struct {
	Name       string    `bson:"name"`
	Category   string    `bson:"category"`
	CreateTime time.Time `bson:"createTime"`
	ModifyTime time.Time `bson:"modifyTime"`
}
