package class_controller

import (
	"mentor/classroom/dto/req"
	"mentor/classroom/dto/res"
	"mentor/classroom/errcode"
	"mentor/classroom/middleware/log"
	"mentor/classroom/service/class_info_service"
	"mentor/classroom/service/class_process_service"
	"mentor/classroom/service/class_time_service"

	"github.com/gin-gonic/gin"
)

func GetLastUnfinishedClass(c *gin.Context) {
	classroomId, _ := c.GetQuery("classroomId")
	studentId, _ := c.GetQuery("studentId")

	err, result := class_info_service.GetLastUnfinishedClassInfo(classroomId, studentId)
	if err != nil {
		switch err.Error() {
		case "no data":
			c.JSON(200, res.NewErrorResponse(errcode.DataNotFindError))
			return
		default:
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	c.JSON(200, res.NewSuccessResponse(result))
	return
}

func InitClass(c *gin.Context) {
	var request req.InitClassReq
	c.BindJSON(&request)

	err, classId := class_process_service.InitClass(request.ClassroomId, request.MentorId, request.StudentId, request.ClassTitle, request.ClassDesc, request.Points, request.ClassTime)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(res.NewInitClassRes(classId)))
	return
}

func ReimburseClass(c *gin.Context) {
	var request req.ReimburseClassReq
	c.BindJSON(&request)

	err := class_process_service.ReimburseClass(request.ClassId)
	if err != nil {
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
		return
	}

	c.JSON(200, res.NewSuccessResponse(nil))
	return
}

func GetStartTime(c *gin.Context) {
	classId, _ := c.GetQuery("classId")

	err, startTime := class_time_service.GetStartTime(classId)
	if err != nil {
		switch err.Error() {
		case "wrong status":
			c.JSON(200, res.NewErrorResponse(errcode.WrongStatusError))
			return
		default:
			log.Logger.Errorf("%+v\n", err)
			c.JSON(200, res.NewErrorResponse(errcode.UnexpectError))
			return
		}
	}

	c.JSON(200, res.NewSuccessResponse(res.GetClassStartTime{StartTime: startTime}))
	return
}

func Panic(c *gin.Context) {
	panic("panic")
}
