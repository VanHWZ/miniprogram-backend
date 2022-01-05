package middleware

import (
	"bytes"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"miniprogram-backend/conf"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	errorLogger = logrus.New()
)

func init() {
	path := filepath.Join(conf.Config.Logger.BaseDir, conf.Config.Logger.ErrorLog)
	if _, err := os.Stat(conf.Config.Logger.BaseDir); err != nil {
		os.MkdirAll(conf.Config.Logger.BaseDir, os.ModePerm)
	}
	errorLogger.SetOutput(&lumberjack.Logger{
		Filename: path,
		MaxSize: 1,
		MaxBackups: 10,
		MaxAge: 180,
		Compress: true,
	})
	errorLogger.SetFormatter(&nested.Formatter{
		HideKeys:    false,
		TimestampFormat: time.RFC3339,
		NoColors: true,
		FieldsOrder: []string{"from", "method", "uri", "status", "cost", "request_header", "request_body", "response_body"},
	})
}

func printStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return fmt.Sprintf("==> %s\n", string(buf[:n]))
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var processTime float64
		var reqMethod, reqUri, reqIP string
		var reqBody []byte
		var reqHeader http.Header
		var startTime, now time.Time

		defer func() {
			if err := recover(); err != nil {
				stack := printStack()
				endTime := time.Now()
				processTime = endTime.Sub(startTime).Seconds()
				entry := errorLogger.WithFields(logrus.Fields{
					"from": reqIP,
					"method": reqMethod,
					"uri": reqUri,
					"request_header": reqHeader,
					"request_body": string(reqBody),
					"status": c.Writer.Status(),
					"cost": fmt.Sprintf("%fs", processTime),
				})
				entry.Time = now
				entry.Error(fmt.Sprintf("\nhttp: panic serving %v: %v\n%s", reqIP, err, stack))
				c.Status(http.StatusInternalServerError)
			}
		}()

		reqMethod = c.Request.Method
		reqUri = c.Request.RequestURI
		reqHeader = c.Request.Header
		reqBody, _ = c.GetRawData()
		reqIP = c.ClientIP()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		startTime = time.Now()
		now = time.Now()

		c.Next()
	}
}