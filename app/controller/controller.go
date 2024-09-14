package controller

import (
	"github.com/giles-wong/roadShow/library/response"
	"github.com/gin-gonic/gin"
)

func HandleAdmin(c *gin.Context) {

	params, _ := c.Get("params")

	response.Success(c, 200, "success", params)
}
