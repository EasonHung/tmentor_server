package controller

import (
	"case_system/config"
	"case_system/dto"
	"case_system/middleware/log"
	token_verify_middleware "case_system/middleware/token_verify"
	"case_system/service/bid_info_service"
	"case_system/service/student_case_service"
	"case_system/service/teacher_case_service"
	"case_system/utils/cloud_storage_utils"
	"case_system/vo"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	studentCaseRoute := route.Group("/studentCase")
	{
		studentCaseRoute.POST("/add", token_verify_middleware.VerifyToken, addStudentCase)
		studentCaseRoute.GET("/", getStudentCase)
		studentCaseRoute.GET("/one", getOneStudentCase)
		studentCaseRoute.GET("/user", token_verify_middleware.VerifyToken, getUserStudenCases)
		studentCaseRoute.POST("/bid/add", token_verify_middleware.VerifyToken, addStudentCaseBid)
		studentCaseRoute.GET("/bids/info", getBidOfStudentCase)
	}

	teacherCaseRoute := route.Group("/teacherCase")
	{
		teacherCaseRoute.POST("/add", token_verify_middleware.VerifyToken, addTeacherCase)
		teacherCaseRoute.GET("/", getTeacherCase)
	}

	bidCaseRoute := route.Group("/bidInfo")
	{
		bidCaseRoute.GET("/user", getUserBidInfo)
	}
}

func getUserBidInfo(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	bidInfos, err := bid_info_service.GetUserBidInfo(userId)
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Logger.Error(fmt.Printf("%+v\n", err))
		c.JSON(500, err)
		return
	}

	getUserBidInfoVo := vo.GetUserBidInfoResVo{}
	getUserBidInfoVo.BidInfoListConvertor(bidInfos)
	c.JSON(200, getUserBidInfoVo)
	return
}

func getUserStudenCases(c *gin.Context) {
	jwtClaim, _ := c.Get("jwtClaim")
	userId := getUserIdFromJwtClaim(jwtClaim)

	studentCaseDtoList, err := student_case_service.GetStudentCasesByUserId(userId)
	if err != nil {
		c.JSON(500, err)
		log.Logger.Error(err)
	}

	getStudentCaseResVo := vo.GetStudentCaseResVo{}
	getStudentCaseResVo.StudentCaseListConvertor(studentCaseDtoList)
	c.JSON(200, getStudentCaseResVo)
	return
}

func getBidOfStudentCase(c *gin.Context) {
	studentCaseId, _ := c.GetQuery("studentCaseId")

	dto, err := student_case_service.GetBidInfos(studentCaseId)
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Logger.Error(err)
		c.JSON(500, err)
		return
	}

	res := vo.GetBidInfoResVo{}
	res.DtoConvertor(dto)

	c.JSON(200, res)
	return
}

func addStudentCaseBid(c *gin.Context) {
	jwtClaim, _ := c.Get("jwtClaim")
	userId := getUserIdFromJwtClaim(jwtClaim)
	var request vo.AddBidReq
	c.BindJSON(&request)

	err := student_case_service.AddBid(userId, request.BidPrice, request.StudentCaseId, request.Classtime)
	if err != nil {
		log.Logger.Error(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, "OK")
}

func addStudentCase(c *gin.Context) {
	form, _ := c.MultipartForm()
	pictureFiles := form.File["picture"]
	pictureUrls := make([]string, 0)
	for _, pictureFile := range pictureFiles {
		pictureUrl, err := cloud_storage_utils.UploadAndGetPublicUrl(config.GLOBAL_CONFIG.Gcp.BucketName, "case_system/student_case/picture/", pictureFile)
		if err != nil {
			log.Logger.Error(err)
			c.JSON(500, err)
			return
		}

		pictureUrls = append(pictureUrls, pictureUrl)
	}

	jwtClaim, _ := c.Get("jwtClaim")
	userId := getUserIdFromJwtClaim(jwtClaim)
	err := student_case_service.AddStudentCase(c.PostForm("avatarUrl"), userId, c.PostForm("nickname"), time.Now(), c.PostForm("title"), c.PostForm("content"), pictureUrls)
	if err != nil {
		log.Logger.Error(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, "ok")
}

func getOneStudentCase(c *gin.Context) {
	studentCaseId, _ := c.GetQuery("studentCaseId")

	studentCase, err := student_case_service.GetStudentCaseByCaseId(studentCaseId)
	if err != nil {
		fmt.Printf("%+v\n", err)
		log.Logger.Error(fmt.Printf("%+v\n", err))
		c.JSON(500, err)
		return
	}
	getStudentCaseResVo := vo.GetStudentCaseResVoItem{}
	getStudentCaseResVo.StudentCaseConvertor(studentCase)
	c.JSON(200, getStudentCaseResVo)
	return
}

func getStudentCase(c *gin.Context) {
	page, _ := c.GetQuery("page")
	pageInt64, _ := strconv.ParseInt(page, 10, 64)

	studentCaseDtoList, err := student_case_service.GetStudentCaseByPage(pageInt64)
	if err != nil {
		c.JSON(500, err)
		log.Logger.Error(err)
	}

	getStudentCaseResVo := vo.GetStudentCaseResVo{}
	getStudentCaseResVo.StudentCaseListConvertor(studentCaseDtoList)
	c.JSON(200, getStudentCaseResVo)
	return
}

func addTeacherCase(c *gin.Context) {
	pictureFile, err := c.FormFile("picture")
	pictureUrl, err := cloud_storage_utils.UploadAndGetPublicUrl(config.GLOBAL_CONFIG.Gcp.BucketName, "case_system/student_case/picture/", pictureFile)
	jwtClaim, _ := c.Get("jwtClaim")

	userId := getUserIdFromJwtClaim(jwtClaim)
	err = teacher_case_service.AddTeacherCase(c.PostForm("avatarUrl"), userId, c.PostForm("nickname"), time.Now(), c.PostForm("title"), c.PostForm("content"), pictureUrl)
	if err != nil {
		log.Logger.Error(err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, "ok")
}

func getTeacherCase(c *gin.Context) {
	page, _ := c.GetQuery("page")
	pageInt64, _ := strconv.ParseInt(page, 10, 64)

	studentCaseDtoList, err := teacher_case_service.GetTeacherCaseByPage(pageInt64)
	if err != nil {
		c.JSON(500, err)
		log.Logger.Error(err)
	}

	getTeacherCaseResVo := vo.GetTeacherCaseResVo{}
	getTeacherCaseResVo.TeacherCaseListConvertor(studentCaseDtoList)
	c.JSON(200, getTeacherCaseResVo)
	return
}

func getUserIdFromJwtClaim(jwtClaim interface{}) string {
	return jwtClaim.(*dto.JwtClaim).UserId
}
