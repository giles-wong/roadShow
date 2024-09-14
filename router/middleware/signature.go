package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
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
		checkSign := md5.Sum([]byte(fmt.Sprintf("%s%s%s", appSecret, stringToBeSigned.String(), appSecret)))
		checkSignStr := fmt.Sprintf("%X", checkSign)
		fmt.Println(stringToBeSigned.String())
		fmt.Println(checkSignStr)
		fmt.Println(signature)

		//timestamp, _ := params["timestamp"].(float64)
		//// 假设时间戳是以毫秒为单位
		//tsInt64 := int64(timestamp)
		//t := time.UnixMilli(tsInt64).UTC().UnixMilli()
		//params["timestamp"] = t
		//
		//// 获取签名
		//signature, ok := params["sign"].(string)
		//if !ok {
		//	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "缺少签名参数"})
		//	return
		//}
		//// 删除签名参数
		//delete(params, "sign")
		//// 构建待签名字符串
		//stringToBeSigned := buildStringToBeSigned(params)
		//appKey := params["app_key"].(string)
		//appSecret, _ := tools.GetAppSecret(appKey)
		//// 生成签名
		//checkSign := generateSignature(stringToBeSigned, appSecret)
		//
		//fmt.Println(checkSign)
		//fmt.Println(signature)
		//
		//// 验证签名
		//if strings.Compare(strings.ToUpper(checkSign), strings.ToUpper(signature)) != 0 {
		//	response.Error(context, 4004, "签名错误")
		//	return
		//}

		// 如果签名正确，则继续处理请求
		context.Next()
	}
}

func stripslashes(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, `\"`, "\""), `\\`, `\`)
}

// 构建签名字符串
func buildStringToBeSigned(params map[string]interface{}) string {
	fmt.Println(params)
	var buffer strings.Builder
	keys := make([]string, 0, len(params))

	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		v := params[k]
		var vStr string
		if isJsonArray(v) {
			jsonStr, _ := json.Marshal(v)
			vStr = string(jsonStr)
		} else if v == true || v == false {
			vStr = boolToString(v.(bool))
		} else {
			if v == "" {
				vStr = ""
			} else {
				vStr = v.(string)
			}
		}
		buffer.WriteString(fmt.Sprintf("%s%s", k, vStr))
	}

	return buffer.String()
}

// generateSignature 生成签名
func generateSignature(data, appSecret string) string {
	fmt.Println(data)
	fmt.Println(appSecret)
	h := md5.New()
	h.Write([]byte(appSecret + data + appSecret))
	return hex.EncodeToString(h.Sum(nil))
}

// boolToString 转换布尔值为字符串
func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// isJsonArray 检查值是否为JSON数组
func isJsonArray(val interface{}) bool {
	if val == nil {
		return false
	}
	switch val.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}
