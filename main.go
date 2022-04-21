package main

import (
	"gin/config"
	"gin/models"
	"gin/routers"
	"log"
	"net/http"
)

func main() {
	//初始化配置
	config.Init()
	//初始化mysql
	models.InitMysql()
	port := config.Viper.GetString("app.port")

	router := routers.InitRoute()
	log.Println("监听端口", "http://127.0.0.1:"+port)
	// 启动并监听服务
	err := http.ListenAndServe(":"+port, router)
	if err != nil {

		return
	}
}
