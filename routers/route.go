package routers

import (
	"gin/config"
	"gin/controller"
	"gin/services/session"
	"gin/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute() *gin.Engine {
	router := gin.New()

	if config.Viper.GetString("debug_mode") == "false" {
		//正式
		router.StaticFS("static", http.FS(static.EmbedStatic))
	} else {
		//测试 避免重启服务器更新静态文件
		router.StaticFS("static", http.Dir("static"))
	}

	sr := router.Group("/", session.EnableCookieSession())
	{
		//sr.GET("/", controller.Index)          //主页
		sr.GET("/login", controller.Login)     //登录
		sr.GET("/register", controller.Logout) //退出

		auth := router.Group("/", session.SessionMiddleware())
		{
			auth.GET("/home", controller.Index) //主页
		}

	}
	return router
}
