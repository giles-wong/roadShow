package admin

import (
	"github.com/giles-wong/roadShow/library/response"
	"github.com/gin-gonic/gin"
)

func HandleAdmin(c *gin.Context) {
	var params map[string]interface{}

	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, 4009, "参数错误")
		return
	}

	response.Success(c, 200, "success", params)
}
