package admin

import (
	"fmt"
	"github.com/giles-wong/roadShow/library/response"
	"github.com/gin-gonic/gin"
)

func HandleAdmin(c *gin.Context) {
	var params map[string]interface{}
	err := c.BindJSON(&params)
	fmt.Println(err)
	if err != nil {
		response.Error(c, 400, "参数绑定失败")
		return
	}

	response.Success(c, 200, "success", params)
}
