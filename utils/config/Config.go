package config

// 配置文件初始化工具

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/giles-wong/roadShow/utils/enum"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App      App      `yaml:"app"`
	Log      Log      `yaml:"log"`
	DbConfig DbConfig `yaml:"dbconf"`
}

// InitConfig 读取yaml配置文件，转换为Config结构体 初始化配置文件
func (config *Config) InitConfig() *Config {
	// 获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 获取环境变量 根据环境变量加载不同的配置
	env := os.Getenv(enum.GoEnv)
	if env == "" {
		env = enum.DefaultEnv
	}
	// 通过 viper 读取配置文件
	vip := viper.New()
	vip.AddConfigPath(path + "/resource") // 设置读取的文件路径
	vip.SetConfigName("config-" + env)    // 设置读取的文件名
	vip.SetConfigType("yaml")             // 设置读取的文件类型
	// 尝试读取配置文件
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	// 监听配置文件
	vip.WatchConfig()
	vip.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了", in.Name)
		if err := vip.Unmarshal(&config); err != nil {
			fmt.Println("配置文件修改了，但是无法重新加载", err)
		}
	})

	err = vip.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}
