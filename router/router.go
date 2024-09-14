package router

import (
	"fmt"
	"github.com/giles-wong/roadShow/app/admin"
	"github.com/giles-wong/roadShow/app/user"
	"github.com/giles-wong/roadShow/global"
	"github.com/giles-wong/roadShow/library/response"
	"github.com/giles-wong/roadShow/router/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil
	}
	/**静态资源处理*/
	//r.Static("/resource", "./resource")
	//访问域名根目录重定向
	r.GET("/", func(c *gin.Context) {
		response.Success(c, 200, "Welcome to roadShow", "")
		//c.Redirect(http.StatusMovedPermanently, global.App.Config.App.Rootview)
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
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// TODO 注册中间件
	// 第一步 必要参数校验
	r.Use(middleware.ValidityAPi())
	// 第二步 签名校验
	r.Use(middleware.Signature())

	r.NoRoute(func(c *gin.Context) {
		fmt.Println(112, c.Request.URL.Path)
		//path := c.Request.URL.Path
		//method := c.Request.Method
		response.Error(c, 404, "请求的接口不存在, 请确认")
	})

	r.POST("/admin", admin.HandleAdmin)
	r.POST("/user", user.HandleAdmin)

	return r
}
