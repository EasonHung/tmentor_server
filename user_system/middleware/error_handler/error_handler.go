package errorhandler

import (
	"mentor_app/user_system/errcode"
	"mentor_app/user_system/middleware/log"
	"mentor_app/user_system/vo"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func MyErrorHandler(c *gin.Context) {
	defer func() {
		// 可以取得 panic 的回傳值
		if r := recover(); r != nil {
			log.Logger.Error("err: ", r, "stacktrace from panic: \n"+string(debug.Stack()))
			// find out exactly what the error was and set err
			switch r.(type) {
			case *errcode.Error:
				err := r.(*errcode.Error)
				c.JSON(200, vo.NewErrorResponse(err))
			default:
				// Fallback err (per specs, error strings should be lowercase w/o punctuation
				c.JSON(200, vo.NewErrorResponse(errcode.UnexpectError))
			}
		}
	}()
	c.Next()
}
