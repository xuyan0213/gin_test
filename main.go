package main

import (
	"gin/config"
	"gin/models"
	"log"
)

func main() {
	config.Init()
	//初始化mysql
	models.Init()
	port := config.Viper.GetString("app.port")

	log.Println("监听端口", "http://127.0.0.1:"+port)
	//// 启动并监听服务
	//err := http.ListenAndServe(":"+port, router)
	//if err != nil {
	//	return
	//}
}
