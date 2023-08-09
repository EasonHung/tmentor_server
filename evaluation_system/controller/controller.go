package controller

import (
	"evaluation_system/errcode"
	"evaluation_system/middleware/log"
	"evaluation_system/service/evaluation_service"
	"evaluation_system/vo"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	evaluationRoute := route.Group("/user")
	{
		evaluationRoute.POST("/new", createNewEvaluation)
		evaluationRoute.GET("/scoreAndCount", getEvaluationScoreAndCount)
	}
	postRoute := route.Group("/post")
	{
		postRoute.POST("/new", newPost)
		postRoute.GET("/user", getPostByUserId)
	}
}

func getEvaluationScoreAndCount(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	count, score, err := evaluation_service.CountUserPostsAndGetAverageScore(userId)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "mongo: no documents in result"):
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.UserNotFoundError))
			return
		default:
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	res := vo.GetScoreCountRes{
		UserId: userId,
		Count:  count,
		Score:  score,
	}
	c.JSON(200, vo.NewSuccessResponse(res))
}

func newPost(c *gin.Context) {
	var newPostReq vo.NewPostReq
	err := c.ShouldBind(&newPostReq)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err.Error())
		return
	}

	err = evaluation_service.NewPost(newPostReq.FromUserId, newPostReq.ToUserId, newPostReq.Score, newPostReq.Description)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "OK")
	return
}

func getPostByUserId(c *gin.Context) {
	userId, _ := c.GetQuery("userId")
	pageString, _ := c.GetQuery("page")
	page, _ := strconv.ParseInt(pageString, 10, 64)

	postsVo, err := evaluation_service.GetUserPosts(page, userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	totalCount, averageScore, err := evaluation_service.CountUserPostsAndGetAverageScore(userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	postsVo.TotalCount = totalCount
	postsVo.AverageScore = averageScore
	c.JSON(200, vo.NewSuccessResponse(postsVo))
	return
}

func createNewEvaluation(c *gin.Context) {
	var createReq vo.CreateEvaluationReq
	err := c.ShouldBind(&createReq)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err.Error())
		return
	}

	err = evaluation_service.CreateNewEvaluation(createReq.UserId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "OK")
	return
}
