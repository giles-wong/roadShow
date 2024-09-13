package middleware

import (
	"fmt"
	"github.com/giles-wong/roadShow/global"
	"github.com/giles-wong/roadShow/library/response"
	"github.com/gin-gonic/gin"
	"strings"
)

var params map[string]interface{}

func ValidityAPi() gin.HandlerFunc {
	return func(context *gin.Context) {
		checkParams := [6]string{"appKey", "method", "timestamp", "timeout", "noncestr", "sign"}
		// 验证必须要存在的参数
		err := context.BindJSON(&params)
		if err != nil {
			response.Error(context, 400, "参数解析失败")
			context.Abort()
			return
		}

		for _, v := range checkParams {
			if params[v] == nil || params[v] == "" {
				response.Error(context, 401, "缺少必要参数:"+v)
				context.Abort()
				return
			}
		}
		// token 验证
		conf := strings.Split(global.App.Config.Signature.NoToken, ",")

		fmt.Println(conf)
		fmt.Println(params["method"])
		if !contains(conf, params["method"].(string)) && params["token"] == nil {
			response.Error(context, 400, "token验证失败01")
			context.Abort()
			return
		}
	}
}

// 判断字符串是否在切片中
func contains(s []string, str string) bool {
	fmt.Println(str)
	fmt.Println(s)
	for _, a := range s {
		if a == str {
			return true
		}
	}
	return false
}
