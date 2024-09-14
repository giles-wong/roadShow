package main

import (
	"fmt"
	"github.com/giles-wong/roadShow/bootstrap"
	"github.com/giles-wong/roadShow/global"
	"runtime"
	"strconv"
)

func main() {
	//初始化配置文件
	global.App.Config.InitConfig()
	// 初始化日志配置
	global.App.Log = bootstrap.InitLog()
	global.App.Log.Info("starting server ... ok")

	// 加载应用配置
	setCpu, _ := strconv.Atoi(global.App.Config.AppConf.CpuNum)
	fmt.Println("setCpu", setCpu)
	machineCpu := runtime.NumCPU()
	if setCpu > machineCpu { // 如果配置cpu核数大于当前计算机核数，则等当前计算机核数
		setCpu = machineCpu
	}

	if setCpu > 0 {
		runtime.GOMAXPROCS(setCpu)
		global.App.Log.Info(fmt.Sprintf("当前计算机核数: %v个,调用：%v个", machineCpu, setCpu))
	} else {
		runtime.GOMAXPROCS(machineCpu)
		global.App.Log.Info(fmt.Sprintf("当前计算机核数: %v个,调用：%v个", machineCpu, setCpu))
	}

	bootstrap.RunServer()
}
