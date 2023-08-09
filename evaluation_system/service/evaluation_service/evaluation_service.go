package evaluation_service

import (
	"context"
	"evaluation_system/dao/db_connection"
	"evaluation_system/dao/evaluation_posts_dao"
	"evaluation_system/dao/user_evaluation_dao"
	"evaluation_system/internalAPI/user_system_api"
	"evaluation_system/vo"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

func CreateNewEvaluation(userId string) error {
	evaluationDao := user_evaluation_dao.UserEvaluation{
		UserId:            userId,
		AverageScore:      0,
		EvaluationPostIds: make([]string, 0),
	}
	_, err := user_evaluation_dao.CreateOne(evaluationDao)
	if err != nil {
		err = errors.Wrap(err, "error create user evaluation")
		return err
	}

	return nil
}

func NewPost(fromUserId string, toUserId string, score int, description string) error {
	transaction := func(ctx context.Context) (interface{}, error) {
		postId := bson.NewObjectId().Hex()
		postDao := evaluation_posts_dao.EvaluationPost{
			PostId:      postId,
			FormUserId:  fromUserId,
			ToUserId:    toUserId,
			Score:       score,
			Description: description,
		}
		_, err := evaluation_posts_dao.InsertOneWithTx(ctx, postDao)
		if err != nil {
			err = errors.Wrap(err, "error insert new posts")
			return nil, err
		}

		err, averageScore := calculateAverageScore(ctx, toUserId, score)
		if err != nil {
			err = errors.Wrap(err, "error calculate average score")
			return nil, err
		}

		err = user_evaluation_dao.UpdateAverageScoreByUserIdWithTx(ctx, toUserId, averageScore)
		if err != nil {
			err = errors.Wrap(err, "error update user average score")
			return nil, err
		}

		err = user_evaluation_dao.PushPostIdByUserIdWithTx(ctx, toUserId, postId)
		if err != nil {
			err = errors.Wrap(err, "error insert new posts")
			return nil, err
		}

		return nil, nil
	}

	_, err := db_connection.MONGO_CLIENT.DoTransaction(context.Background(), transaction)

	return err
}

func GetUserPosts(page int64, userId string) (vo.GetUserPostRes, error) {
	resVo := vo.GetUserPostRes{}
	posts, err := evaluation_posts_dao.FindByToUserWithPage(page, userId)
	if err != nil {
		err = errors.Wrap(err, "error get user posts")
		return resVo, err
	}
	resVo.EvaluationPostDaoConvertor(posts)

	for index, post := range resVo.Posts {
		err, userInfo := user_system_api.GetUserInfo(post.FromUserId)
		if err != nil {
			err = errors.Wrap(err, "error get user info")
			return resVo, err
		}

		// 用index 才可以換掉object裡面的member
		resVo.Posts[index].FromUserAvatarUrl = userInfo.Data.AvatorUrl
		resVo.Posts[index].FromUserNickname = userInfo.Data.Nickname
		resVo.Posts[index].FromUserProfession = userInfo.Data.Profession[0]
	}

	return resVo, nil
}

func CountUserPostsAndGetAverageScore(userId string) (int, float64, error) {
	userEvaluation, err := user_evaluation_dao.FindByUserId(userId)
	if err != nil {
		err = errors.Wrap(err, "error get user evaluation")
		return 0, 0, err
	}

	return len(userEvaluation.EvaluationPostIds), userEvaluation.AverageScore, nil
}

func calculateAverageScore(ctx context.Context, userId string, score int) (error, float64) {
	userEvaluation, err := user_evaluation_dao.FindByUserIdWithTx(ctx, userId)
	if err != nil {
		err = errors.Wrap(err, "error get user evaluation")
		return err, 0
	}

	postCount := len(userEvaluation.EvaluationPostIds)
	return nil, (userEvaluation.AverageScore*float64(postCount) + float64(score)) / (float64(postCount) + 1)
}
