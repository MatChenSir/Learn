package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SaveLog() {
	fmt.Println("Saving log")
	r := gin.Default()

	// 使用自定义的中间件函数来记录日志
	r.Use(Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}

//结合第三方库如lumberjack来实现日志文件滚动。
func setLogConfig() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   getCurrentFilePath() + "/logs/requestAndResponse",
		MaxSize:    500,   // 日志文件最大大小（MB）
		MaxBackups: 10,    // 最大保留旧日志文件的数量
		MaxAge:     7,     // 最大保留旧日志文件的天数
		Compress:   false, // 是否压缩旧日志文件
	}

}

// getCurrentFile 获取项目目录
func getCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	absPath, _ := filepath.Abs(file)
	//有几层可以嵌套几层filepath.Dir，也可以通过别的方式定义货找到路径  filepath.Dir(filepath.Dir(absPath))
	return filepath.Dir(absPath)

}

func getActiontFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	absPath, _ := filepath.Abs(file)
	_, _, line, _ := runtime.Caller(1)
	return absPath + strconv.Itoa(line)
}

type CustomFieldsFormatter struct {
	logrus.TextFormatter
}

// Format 格式化日志条目
func (f *CustomFieldsFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 添加额外的字段
	entry.Data["RequestTime"] = time.Now().Format(time.RFC3339)
	entry.Data["ActionFile"] = getActiontFilePath()
	return f.TextFormatter.Format(entry)
}

const MaxBodyLength = 2000 // 设置最大Body长度为2000
var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

func Logger() gin.HandlerFunc {
	//setupLogging(duration)

	lumberjackLogger := setLogConfig()
	// 创建一个新的logrus实例
	log := logrus.New()

	// 创建自定义的 TextFormatter
	formatter := &CustomFieldsFormatter{
		TextFormatter: logrus.TextFormatter{
			DisableColors:    false,        // 启用颜色输出
			ForceColors:      true,         // 启用强制颜色输出
			DisableTimestamp: false,        // 启用时间戳
			TimestampFormat:  time.RFC3339, // 时间戳格式
		},
	}

	// 配置日志输出为JSON格式
	//log.Formatter = &logrus.JSONFormatter{}
	log.SetFormatter(formatter)

	// 将日志输出到lumberjackLogger
	log.Out = lumberjackLogger
	log.SetLevel(logrus.InfoLevel)

	// 输出日志
	return func(c *gin.Context) {
		t := time.Now()
		requestData := ""
		responseData := ""
		//记录请求body
		queryParams := c.Request.URL.Query()
		requestData = fmt.Sprintf(" request params : %s;", queryParams)
		requestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Errorf("[GIN]: Error reading request body")
		} else {
			// 重置 body，以便后续的处理
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
			if len(requestBody) > MaxBodyLength {
				requestBody = []byte("Body too long, size: " + strconv.Itoa(len(requestBody)))
			}
			requestData = requestData + fmt.Sprintf(" request body :%v;", string(requestBody))
		}

		// process request
		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		methodColor := colorForMethod(method)
		path := c.Request.URL.Path
		if res, ok := c.Get("response"); ok {
			//responseData = fmt.Sprintf("%v", res.(string))
			jsonData, err := json.Marshal(res)
			if err != nil {
				responseData = "parse error: " + err.Error()
			} else {
				responseData = string(jsonData)
			}

		}
		//截断或省略响应的Body
		if len(responseData) > MaxBodyLength {
			responseData = "Body too long, size: " + strconv.Itoa(len(responseData))
		}
		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				log.Warningf("[GIN] |当前返回状态: %s %3d %s|启动耗费时长 %12v |IP地址: %s |请求方式为:%s %s %s; 请求路径: %s ;报错信息:%-7s;  | request: %s |response:%s|\n",
					statusColor,
					statusCode,
					reset,
					latency,
					clientIP,
					methodColor, method, reset,
					path,
					c.Errors.String(),
					requestData,
					responseData,
				)
			}
		case statusCode >= 500:
			{
				log.Errorf("[GIN] |当前返回状态: %s %3d %s|启动耗费时长 %12v |IP地址: %s |请求方式为:%s %s %s; 请求路径: %s ;报错信息:%-7s;  | request: %s |response:%s|\n",
					statusColor,
					statusCode,
					reset,
					latency,
					clientIP,
					methodColor, method, reset,
					path,
					c.Errors.String(),
					requestData,
					responseData,
				)
			}
		default:
			log.Infof("[GIN] |当前返回状态: %s %3d %s|启动耗费时长 %12v |IP地址: %s |请求方式为:%s %s %s; 请求路径: %s ;报错信息:%-7s;  | request: %s |response:%s|\n",
				statusColor,
				statusCode,
				reset,
				latency,
				clientIP,
				methodColor, method, reset,
				path,
				c.Errors.String(),
				requestData,
				responseData,
			)
		}

	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code <= 299:
		return green
	case code >= 300 && code <= 399:
		return white
	case code >= 400 && code <= 499:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return blue
	case method == "POST":
		return cyan
	case method == "PUT":
		return yellow
	case method == "DELETE":
		return red
	case method == "PATCH":
		return green
	case method == "HEAD":
		return magenta
	case method == "OPTIONS":
		return white
	default:
		return reset
	}
}
