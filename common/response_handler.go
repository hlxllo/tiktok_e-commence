package common

import (
	"github.com/gin-gonic/gin"
)

// Response 通用响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HandleResponse 统一响应返回
func HandleResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Code:    status,
		Message: message,
		Data:    data,
	})
}
