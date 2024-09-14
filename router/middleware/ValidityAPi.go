package middleware

import (
	"github.com/giles-wong/roadShow/library/response"
	"github.com/giles-wong/roadShow/utils/tools"
	"github.com/gin-gonic/gin"
)

var params map[string]interface{}

func ValidityAPi() gin.HandlerFunc {
	return func(context *gin.Context) {
		checkParams := [...]string{"app_key", "app_version", "method", "timestamp", "timeout", "noncestr", "sign"}
		// 验证必须要存在的参数
		err := context.BindJSON(&params)
		if err != nil {
			response.Error(context, 4001, "参数解析失败")
			context.Abort()
			return
		}

		for _, v := range checkParams {
			if params[v] == nil || params[v] == "" {
				response.Error(context, 4002, "缺少必要参数:"+v)
				context.Abort()
				return
			}
		}
		// 验证平台是否合法
		_, err = tools.GetAppSecret(params["app_key"].(string))
		if err != nil {
			response.Error(context, 4003, "不合法的平台，请确认")
			context.Abort()
			return
		}
	}
}
