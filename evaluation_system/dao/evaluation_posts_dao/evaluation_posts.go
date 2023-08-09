package evaluation_posts_dao

import "time"

type EvaluationPost struct {
	PostId      string    `bson:"postId"`
	FormUserId  string    `bson:"fromUserId"`
	ToUserId    string    `bson:"toUserId"`
	Score       int       `bson:"score"`
	Description string    `bson:"description"`
	CreateTime  time.Time `bson:"createTime"`
	ModifyTime  time.Time `bson:"modifyTime"`
}
