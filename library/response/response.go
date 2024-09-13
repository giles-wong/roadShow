package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Error 返回错误信息
func Error(code int, message string, c *gin.Context) {
	response := errorResponse{Code: code, Message: message}

	c.JSON(http.StatusOK, response)
}

// Success 返回成功信息
func Success(code int, message string, data interface{}, c *gin.Context) {

	response := successResponse{Code: code, Message: message, Data: data}

	c.JSON(http.StatusOK, response)
}
