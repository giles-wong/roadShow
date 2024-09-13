package bootstrap

import (
	"fmt"
	"github.com/giles-wong/roadShow/global"
	"github.com/giles-wong/roadShow/router"
)

// RunServer 优雅重启/停止服务器
func RunServer() {
	//加载路由
	r := router.InitRouter()
	fmt.Println("router:", r.Routes())
	r.Run(":" + global.App.Config.App.Port)
}
