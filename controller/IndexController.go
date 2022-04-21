package controller

import (
	"gin/services/userService"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录
func Login(c *gin.Context) {
	//TODO 登录逻辑 放到服务层
	userService.Login(c)
}

// Logout 退出
func Logout(c *gin.Context) {
	//todo 退出
	userService.Logout(c)
}

// Index 主页
func Index(c *gin.Context) {
	//1. 判断用户是否登录
	userInfo := userService.GetUserInfo(c)
	//已经登录重定向到主页
	if len(userInfo) > 0 {
		c.Redirect(http.StatusFound, "home.html")
	}
	//没有登陆显示登录页面
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
