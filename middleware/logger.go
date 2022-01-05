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
	"miniprogram-backend/util"
	"os"
	"path/filepath"
	"time"
)

var (
	httpLogger = logrus.New()
)

func init() {
	path := filepath.Join(conf.Config.Logger.BaseDir, conf.Config.Logger.HttpLog)
	if _, err := os.Stat(conf.Config.Logger.BaseDir); err != nil {
		os.MkdirAll(conf.Config.Logger.BaseDir, os.ModePerm)
	}
	httpLogger.SetOutput(&lumberjack.Logger{
		Filename: path,
		MaxSize: 1,
		MaxBackups: 10,
		MaxAge: 180,
		Compress: true,
	})
	httpLogger.SetFormatter(&nested.Formatter{
		HideKeys:    false,
		TimestampFormat: time.RFC3339,
		NoColors: true,
		FieldsOrder: []string{"from", "method", "uri", "status", "cost", "request_header", "request_body", "response_body"},
	})
}

func HttpLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &util.BodyLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		reqHeader := c.Request.Header
		reqBody, _ := c.GetRawData()
		reqIP := c.ClientIP()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		startTime := time.Now()
		now := time.Now()

		c.Next()

		endTime := time.Now()
		processTime := endTime.Sub(startTime).Seconds()
		statusCode := c.Writer.Status()
		repBody := blw.Body.String()

		entry := httpLogger.WithFields(logrus.Fields{
			"from": reqIP,
			"method": reqMethod,
			"uri": reqUri,
			"request_header": reqHeader,
			"request_body": string(reqBody),
			"status": statusCode,
			"cost": fmt.Sprintf("%fs", processTime),
			"response_body": repBody,
		})
		entry.Time = now
		entry.Info("http request log")
	}
}