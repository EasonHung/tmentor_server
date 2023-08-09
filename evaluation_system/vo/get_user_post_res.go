package vo

import (
	"evaluation_system/dao/evaluation_posts_dao"
	"time"
)

type GetUserPostRes struct {
	TotalCount   int        `json:"totalCount"`
	AverageScore float64    `json:"averageScore"`
	Posts        []UserPost `json:"posts"`
}

type UserPost struct {
	PostId             string    `json:"postId"`
	FromUserId         string    `json:"fromUserId"`
	FromUserAvatarUrl  string    `json:"fromUserAvatar"`
	FromUserNickname   string    `json:"fromUserNickname"`
	FromUserProfession string    `json:"fromUserProfession"`
	Score              int       `json:"score"`
	Description        string    `json:"description"`
	CreateTime         time.Time `json:"createTime"`
}

func (this *GetUserPostRes) EvaluationPostDaoConvertor(evaluationPostDaos []evaluation_posts_dao.EvaluationPost) {
	postRes := make([]UserPost, 0)
	for _, post := range evaluationPostDaos {
		userPost := UserPost{
			PostId:      post.PostId,
			FromUserId:  post.FormUserId,
			Score:       post.Score,
			Description: post.Description,
			CreateTime:  post.CreateTime,
		}
		postRes = append(postRes, userPost)
	}

	this.Posts = postRes
}
