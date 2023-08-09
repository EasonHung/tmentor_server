package info_controller

import (
	"mentor/classroom/domain/classroom"
	"mentor/classroom/dto/req"
	"mentor/classroom/dto/res"
	"mentor/classroom/errcode"
	"mentor/classroom/internal_error"
	"mentor/classroom/middleware/log"
	"mentor/classroom/service/class_record_service"
	"mentor/classroom/service/class_setting_service"
	"mentor/classroom/service/classroom_info_service"
	"mentor/classroom/service/classroom_registry_service"
	"mentor/classroom/service/init_service"
	"mentor/classroom/service/list_service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetClassroomToken(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	err, token := classroom_registry_service.GetUserClassroomToken(userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(token))
	return
}

func GetClassroomList(c *gin.Context) {
	userId, _ := c.Get("userId")
	classroomInfoList, err := list_service.GetClassroomList(userId.(string))
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}
	c.JSON(200, res.NewSuccessResponse(classroomInfoList))
}

func EnrollClassroom(c *gin.Context) {
	var request req.EnrollReq
	c.BindJSON(&request)

	err := classroom_registry_service.EnrollClassroom(request.ClassroomToken, request.UserId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(nil))
	return
}

func InitUser(c *gin.Context) {
	var request req.InitClassroomReq
	c.BindJSON(&request)

	err := init_service.InitUser(request.UserId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(nil))
}

func AddClassSetting(c *gin.Context) {
	var request req.AddClassSettingReq
	userId, _ := c.Get("userId")
	c.BindJSON(&request)

	newClassSetting := classroom.NewClassSettingFromAddClassSettingReq(request)
	err := class_setting_service.AddClassSetting(userId.(string), newClassSetting)
	if err != nil {
		if errors.As(err, &internal_error.OutOfLimitError{}) {
			log.Logger.Errorf("%+v\n", err)
			res := res.NewErrorResponse(errcode.OutOfLimitError)
			c.JSON(200, res)
			return
		} else if errors.As(err, &internal_error.DuplicateKeyError{}) {
			log.Logger.Errorf("%+v\n", err)
			res := res.NewErrorResponse(errcode.DuplicateKeyError)
			c.JSON(200, res)
			return
		} else {
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	c.JSON(200, res.NewSuccessResponse(nil))
}

func UpdateClassSetting(c *gin.Context) {
	var request req.UpdateClassSettingReq
	c.BindJSON(&request)
	userId, _ := c.Get("userId")

	updatedClassSetting := classroom.NewClassSettingFromUpdateClassSettingReq(request)
	err := class_setting_service.UpdateClassSetting(userId.(string), updatedClassSetting)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(nil))
}

func DeleteClassSetting(c *gin.Context) {
	var request req.DeleteClassSettingReq
	c.BindJSON(&request)
	userId, _ := c.Get("userId")

	err := class_setting_service.DeleteClassSetting(userId.(string), request.SettingName)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(nil))
}

func GetClassSetting(c *gin.Context) {
	userId, _ := c.Get("userId")
	classSettingRes, err := class_setting_service.GetClassSettingList(userId.(string))
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}
	c.JSON(200, res.NewSuccessResponse(classSettingRes))
}

func GetClassroomStatus(c *gin.Context) {
	classroomId, _ := c.GetQuery("classroomId")

	classroomStatusRes, err := classroom_info_service.GetClassroomStatus(classroomId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}
	c.JSON(200, res.NewSuccessResponse(classroomStatusRes))
	return
}

func GetUserClassroom(c *gin.Context) {
	userId, _ := c.Get("userId")

	classroomId, err := classroom_info_service.GetClassroomId(userId.(string))
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(500, err)
		return
	}

	classroomInfoRes := res.GetClassroomOwnershipRes{
		ClassroomId: classroomId,
		UserId:      userId.(string),
	}
	c.JSON(200, classroomInfoRes)
	return
}

func GetUserStudentCount(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	count, err := classroom_info_service.GetStudentCount(userId)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "no documents in result"):
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, res.NewErrorResponse(errcode.UserNotFoundError))
			return
		default:
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	result := res.GetStudentCountRes{
		UserId: userId,
		Count:  count,
	}
	c.JSON(200, res.NewSuccessResponse(result))
	return
}

func GetUserClassroomId(c *gin.Context) {
	userId, _ := c.GetQuery("userId")

	userClassroomId, err := classroom_info_service.GetClassroomId(userId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(userClassroomId))
	return
}

func GetClassRecord(c *gin.Context) {
	userId, _ := c.Get("userId")

	err, record := class_record_service.GetClassRecord(userId.(string))
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(record.ClassRecordList))
	return
}
