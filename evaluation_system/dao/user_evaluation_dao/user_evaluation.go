package user_evaluation_dao

import "time"

type UserEvaluation struct {
	UserId            string    `bson:"userId"`
	EvaluationPostIds []string  `bson:"evaluationPostIds"`
	AverageScore      float64   `bson:"averageScore"`
	CreateTime        time.Time `bson:"createTime"`
	ModifyTime        time.Time `bson:"modifyTime"`
}
