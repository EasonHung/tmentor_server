package log

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = NewLogger()
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func NewLogger() *logrus.Logger {
	// 例項化
	logger := logrus.New()
	// 設定輸出
	// mw := io.MultiWriter(os.Stdout, src) => multiple output resource
	// logger.Out = src

	// 設定日誌級別
	logger.SetLevel(logrus.DebugLevel)

	//設定日誌格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}

func LoggerMiddleware() gin.HandlerFunc {

	logger := NewLogger()

	return func(c *gin.Context) {
		// 開始時間
		startTime := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 處理請求
		c.Next()

		// 結束時間
		endTime := time.Now()

		// 執行時間
		latencyTime := endTime.Sub(startTime)

		// 請求方式
		reqMethod := c.Request.Method

		// 請求路由
		reqUri := c.Request.RequestURI

		// trace id
		traceId := c.GetHeader("request-id")

		// 狀態碼
		statusCode := c.Writer.Status()

		respBody := blw.body.String()

		// 請求IP
		clientIP := c.ClientIP()

		// 日誌格式
		logger.Infof("| %3d | %13v | %15s | %s | %s | %s | %s |",
			statusCode,
			latencyTime,
			traceId,
			clientIP,
			reqMethod,
			reqUri,
			respBody,
		)
	}
}
