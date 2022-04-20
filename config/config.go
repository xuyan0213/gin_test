package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Viper *viper.Viper

func Init() {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}

	// 初始化配置文件
	Viper = viper.New()
	Viper.AddConfigPath(path + "/config")
	Viper.SetConfigName("config")
	Viper.SetConfigType("yaml")
	fmt.Printf("%s\n", path)
	err = Viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败

		// Config file was found but another error was produced
		log.Println("read config error")
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}
