package controller

import (
	"mentor_app/user_system/dto"
	"mentor_app/user_system/errcode"
	"mentor_app/user_system/mentor_redis"
	"mentor_app/user_system/middleware/log"
	token_verify_middleware "mentor_app/user_system/middleware/token_verify"
	"mentor_app/user_system/service/backstage_service"
	"mentor_app/user_system/service/token_verification_service"
	"mentor_app/user_system/service/user_service"
	"mentor_app/user_system/vo"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Route(route *gin.Engine) {
	loginRoute := route.Group("/user")
	{
		loginRoute.POST("/thirdParty", CreateUserWithThirdParty)
		loginRoute.POST("/thirdParty/login", ThirdPartyLogin)
		loginRoute.POST("/token/refresh", RefreshToken)
		loginRoute.POST("/token/verify", VerifyTokenAndReturnUserInfo) // inner api!!
		// loginRoute.GET("/walletId", GetWalletIdByUserId)               // inner api!!
		loginRoute.GET("/cache", GetCache) // inner api!!
	}
	userInfoRoute := route.Group("/userInfo")
	{
		userInfoRoute.GET("/", GetUserInfo)
		userInfoRoute.GET("/avatorUrl", GetUserAvatorUrl)
		userInfoRoute.GET("/cards", GetCards)
		userInfoRoute.POST("/avatar/update", token_verify_middleware.VerifyToken, UpdateAvator)
		userInfoRoute.POST("/picture/update", token_verify_middleware.VerifyToken, UpdatePicture)
		userInfoRoute.GET("/fcmTokenAndNickname", GetFcmTokenAndNickname)
		userInfoRoute.POST("/fcmToken/insert", InsertFcmToken)
		userInfoRoute.POST("/fcmToken/delete", DeleteFcmToken)
		userInfoRoute.POST("/fcmToken/update", UpdateFcmToken)
		userInfoRoute.POST("/update", token_verify_middleware.VerifyToken, UpdateUserInfo)
	}
	backstageRoute := route.Group("/backstage")
	{
		backstageRoute.GET("/token/blackList", GetTokenBlackList)            // inner api!!
		backstageRoute.POST("/token/blackList/add", AddTokenBlackList)       // inner api!!
		backstageRoute.POST("/token/blackList/delete", DeleteTokenBlackList) // inner api!!
		backstageRoute.POST("/professionCatelog/update", UpdateProfessionCatelog)
		backstageRoute.POST("/professionCatelog/add", AddProfessionCategory)
		backstageRoute.GET("/professionCatelog/categoryList", GetCategoryList)
		backstageRoute.GET("/professionCatelog/unCategorizeName", GetUnCategorizeName)
	}
}

func GetTokenBlackList(c *gin.Context) {
	err, list := backstage_service.GetTokenBlackList()
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(list))
	return
}

func DeleteTokenBlackList(c *gin.Context) {
	var request vo.DeleteTokenBlackListReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err = backstage_service.DeleteUserIdIntoBlackList(request.UserId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func AddTokenBlackList(c *gin.Context) {
	var request vo.AddTokenBlackListReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err = backstage_service.AddUserIdIntoBlackList(request.UserId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func VerifyTokenAndReturnUserInfo(c *gin.Context) {
	var request vo.VerifyTokenReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err, claim := token_verification_service.VerifyTokenWithApi(request.UserToken)
	if err != nil {
		switch {
		case err.Error() == "verify fail":
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.TokenDecodeError))
			return
		case err.Error() == "token expired":
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.TokenExpiredError))
			return
		default:
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	userId := claim.UserId
	userDao, err := user_service.GetUser(userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	user := vo.VerifyTokenRes{}
	user.UserDaoConvertor(userDao)

	c.JSON(200, vo.NewSuccessResponse(user))
}

func GetUnCategorizeName(c *gin.Context) {
	res, err := backstage_service.GetUnCategorizeNames()
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, res)
	return
}

func GetCategoryList(c *gin.Context) {
	res, err := backstage_service.GetProfessionCategoryList()
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, res)
	return
}

func AddProfessionCategory(c *gin.Context) {
	var request vo.AddProfssionCatelogReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err = backstage_service.AddProfessionCategory(request.Name, request.Category)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
	return
}

func UpdateProfessionCatelog(c *gin.Context) {
	var request vo.UpdateProfssionCatelogReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err = backstage_service.UpdateProfessionCatelog(request.Name, request.Category, request.OldCategory)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, "ok")
	return
}

func UpdateAvator(c *gin.Context) {
	jwtClaim, _ := c.Get("jwtClaim")
	userId := getUserIdFromJwtClaim(jwtClaim)
	avatorFile, err := c.FormFile("userAvator")
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	err = user_service.UpdateUserAvator(userId, avatorFile)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func GetFcmTokenAndNickname(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	res := vo.GetFcmTokenAndNicknameRes{}
	pushNotificationInfoDto, err := user_service.GetFcmTokenAndNickname(userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}
	res.DtoConvertor(pushNotificationInfoDto)
	c.JSON(200, vo.NewSuccessResponse(res))
}

func InsertFcmToken(c *gin.Context) {
	var request vo.InsertFcmtokenReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err = user_service.InsertFcmToken(request.UserId, request.FcmToken)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}
	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func DeleteFcmToken(c *gin.Context) {
	var request vo.DeleteFcmtokenReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err = user_service.DeleteFcmToken(request.UserId, request.FcmToken)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}
	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func UpdateFcmToken(c *gin.Context) {
	var request vo.UpdateFcmtokenReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	err = user_service.UpdateFcmToken(request.UserId, request.OriginFcmToken, request.NewFcmToken)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}
	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func GetCards(c *gin.Context) {
	pageString, _ := c.GetQuery("page")
	gender, _ := c.GetQueryArray("gender")
	fields, _ := c.GetQueryArray("fields")
	page, _ := strconv.ParseInt(pageString, 10, 64)

	cards, err := user_service.GetCardsByPage(c.GetHeader("request_id"), page, gender, fields)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}
	c.JSON(200, vo.NewSuccessResponse(cards))
	return
}

func UpdatePicture(c *gin.Context) {
	jwtClaim, _ := c.Get("jwtClaim")
	userId := getUserIdFromJwtClaim(jwtClaim)
	pictureFile, err := c.FormFile("userPicture")
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return	
	}

	err = user_service.UpdateUserPicture(userId, pictureFile)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func UpdateUserInfo(c *gin.Context) {
	jwtClaim, _ := c.Get("jwtClaim")
	var request vo.UpdateUserInfoReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	userId := getUserIdFromJwtClaim(jwtClaim)

	request.ValidationAndRebuildRequest()
	userInfoEntity := request.ToUserInfoDao()

	err = user_service.UpdateUserInfo(userId, userInfoEntity, request.UserStatus)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, vo.NewSuccessResponse(nil))
	return
}

func GetUserAvatorUrl(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	userInfo, err := user_service.GetUserInfo(userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, userInfo.AvatorUrl)
}

func GetUserInfo(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	userInfo, err := user_service.GetUserInfo(userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	res := vo.GetUserInfoRes{}
	res.UserInfoDaoConvertor(userInfo)
	c.JSON(200, vo.NewSuccessResponse(res))
}

func CreateUserWithThirdParty(c *gin.Context) {
	var request vo.CreateThirdPartyUserReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	dto := dto.NewUserDTO(request.ThirdPartyId, request.ThirdPartyInfo)
	dto, err = user_service.CreateUserWithThirdParty(dto, c.GetHeader("request_id"))
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
		return
	}

	response := vo.NewCreateUserRes(dto.UserId)
	c.JSON(200, vo.NewSuccessResponse(response))
}

func ThirdPartyLogin(c *gin.Context) {
	var request vo.LoginWithThirdParyReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	accessToken, refreshToken, userId, userStatus, err := user_service.UserLoginWithThirdParty(request.ThirdPartyId, request.ThirdPartyAccessToken)
	if err != nil && err.Error() == "mongo: no documents in result" {
		c.JSON(401, "無此帳號")
		return
	}
	if err != nil {
		c.JSON(500, err)
		return
	}

	response := vo.NewLoginRes(accessToken, refreshToken, userId, userStatus)
	c.JSON(200, vo.NewSuccessResponse(response))
}

func RefreshToken(c *gin.Context) {
	var request vo.RefreshTokenReq
	err := c.BindJSON(&request)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, vo.NewErrorResponse(errcode.ReqBodyError))
		return
	}

	newAccessToken, newRefreshToken, err := user_service.RefreshToken(request.UserToken)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "token decode error"):
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.TokenDecodeError))
			return
		case strings.Contains(err.Error(), "userToken is in the black list"):
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.UserBlackListError))
			return
		case strings.Contains(err.Error(), "token expired"):
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.TokenExpiredError))
			return
		default:
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	response := vo.NewRefreshTokenRes(newAccessToken, newRefreshToken)
	c.JSON(200, vo.NewSuccessResponse(response))
	return
}

func GetCache(c *gin.Context) {
	val, _ := c.GetQuery("key")
	res := mentor_redis.Get(val)

	c.JSON(200, res)
}

func getUserIdFromJwtClaim(jwtClaim interface{}) string {
	return jwtClaim.(*dto.JwtClaim).UserId
}
