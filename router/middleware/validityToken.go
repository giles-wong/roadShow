package middleware

import (
	"fmt"
	"github.com/giles-wong/roadShow/global"
	"github.com/giles-wong/roadShow/utils/tools"
	"github.com/gin-gonic/gin"
	"strings"
)

// ValidityToken 验证token
func ValidityToken() gin.HandlerFunc {
	return func(context *gin.Context) {

		noTokenConf := tools.Slice{
			Slice: strings.Split(global.App.Config.SignConf.NoToken, ","),
		}
		// 不在列表中则校验token
		if noTokenConf.SliceContains(params["method"].(string)) == false {
			token := params["token"]
			//去 redis 中查询用户信息 验证token 是否有效 TODO
			fmt.Println(token)
		}

		context.Next()
	}
}
