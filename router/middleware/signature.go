package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/giles-wong/roadShow/library/response"
	"github.com/giles-wong/roadShow/utils/tools"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
	"time"
)

func Signature() gin.HandlerFunc {
	return func(context *gin.Context) {
		signature, _ := params["sign"].(string)

		paramsChk := make(map[string]interface{})
		for k, v := range params {
			if k != "sign" {
				paramsChk[k] = v
			}
		}
		keys := make([]string, 0, len(params))
		for k := range paramsChk {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var stringToBeSigned strings.Builder
		for _, k := range keys {
			v := params[k]
			var vStr string
			if k == "timestamp" {
				timestamp, _ := params["timestamp"].(float64)
				// 假设时间戳是以毫秒为单位
				tsInt64 := int64(timestamp)
				t := time.UnixMilli(tsInt64).UTC().UnixMilli()
				v = t
			}

			switch v := v.(type) {
			case string:
				if !strings.HasPrefix(v, "@") {
					vStr = stripslashes(v)
				} else {
					vStr = v
				}
			case bool:
				vStr = fmt.Sprintf("%t", v)
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
				vStr = fmt.Sprintf("%v", v)
			case []interface{}:
				vBytes, _ := json.Marshal(v)
				vStr = string(vBytes)
			default:
				// 对于其他类型，直接跳过或处理
				continue
			}

			stringToBeSigned.WriteString(fmt.Sprintf("%s%s", k, vStr))
		}
		appKey := params["app_key"].(string)
		appSecret, _ := tools.GetAppSecret(appKey)

		input := fmt.Sprintf("%s%s%s", appSecret, stringToBeSigned.String(), appSecret)
		// 创建一个新的 MD5 计算器
		hasher := md5.New()
		_, err := hasher.Write([]byte(input))
		if err != nil {
			response.Error(context, 4006, "签名错误")
		}
		// 获取计算后的散列值
		hashBytes := hasher.Sum(nil)
		// 将字节切片转换为十六进制格式的字符串
		checkSign := hex.EncodeToString(hashBytes)

		if strings.Compare(strings.ToUpper(checkSign), strings.ToUpper(signature)) != 0 {
			response.Error(context, 4008, "签名错误02")
			return
		}

		context.Next()
	}
}

func stripslashes(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, `\"`, "\""), `\\`, `\`)
}
