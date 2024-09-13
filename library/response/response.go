package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type errorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
}

type successResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
}

// Error 返回错误信息
func Error(c *gin.Context, code int, message string) {
	response := errorResponse{
		Code:    code,
		Message: message,
		Data:    "",
		Time:    time.Now().Unix(),
	}

	c.JSON(http.StatusOK, response)
}

// Success 返回成功信息
func Success(c *gin.Context, code int, message string, data interface{}) {

	response := successResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Time:    time.Now().Unix(),
	}

	c.JSON(http.StatusOK, response)
}
