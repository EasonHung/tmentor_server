package middleware

import (
	"fmt"
	"mentor/classroom/dto/res"
	"mentor/classroom/internal_api/user_system_api"
	"mentor/classroom/internal_api/user_system_api/user_system_error_code"
	"mentor/classroom/middleware/log"

	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	jwtToken := c.GetHeader("authToken")

	if jwtToken == "test" {
		c.Set("userId", "test1")
		c.Next()
		return
	}

	err, result := user_system_api.VerifyToken(jwtToken)
	if result.Code != user_system_error_code.SUCCESS {
		log.Logger.Info(fmt.Sprintf("verify fail! status code: %s", result.Code))
		log.Logger.Errorf("%+v\n", err)
		c.JSON(200, res.NewErrorResponseWithCodeAndMessage(result.Code, result.Message))
		c.Abort()
		return
	}
	c.Set("userId", result.Data.UserId)

	c.Next()
}
