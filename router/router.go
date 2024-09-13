package router

import (
	"github.com/giles-wong/roadShow/global"
	"github.com/giles-wong/roadShow/library/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	/**静态资源处理*/
	r.Static("/resource", "./resource")
	//访问域名根目录重定向
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, global.App.Config.App.Rootview)
	})

	//控制台日志级别
	gin.SetMode(global.App.Config.App.RunlogType)
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	//跨域访问-注意跨域要放在gin.Default下
	var strArr []string
	if global.App.Config.App.Allowurl != "" {
		strArr = strings.Split(global.App.Config.App.Allowurl, `,`)
	} else {
		strArr = []string{"http://localhost:8080"}
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins: strArr,
		// AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "Businessid", "verify-encrypt", "ignoreCancelToken", "verify-time"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// TODO 注册中间件

	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		response.Error(404, path+" Not Found "+method, c)
	})

	return r
}
