package userService

import (
	"gin/models"
	"gin/services/helper"
	"gin/services/session"
	"gin/services/validator"
	"github.com/gin-gonic/gin"
	"strconv"

	"net/http"
)

func Login(c *gin.Context) {
	//接收参数
	username := c.PostForm("username")
	password := c.PostForm("password")
	avatarId := c.PostForm("avatar_id")

	//验证
	var u validator.User

	u.Username = username
	u.Password = password
	u.AvatarId = avatarId

	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  err.Error(),
		})
	}

	userInfo := models.GetUserByField("username", username)
	md5Pwd := helper.Md5Encrypt(password)

	if userInfo.ID > 0 {
		//用户已存在 对比密码
		if userInfo.Password != md5Pwd {
			//密码错误
			c.JSON(http.StatusOK, gin.H{
				"code": 10002,
				"msg":  "密码错误",
			})
		}
		//更新用户的头像
		models.UpdateAvatar(avatarId, userInfo)
	} else {
		//新用户
		userInfo = models.AddUser(map[string]interface{}{
			"username":  username,
			"password":  md5Pwd,
			"avatar_id": avatarId,
		})
	}

	if userInfo.ID > 0 {
		//保存用户session
		session.SaveUserSession(c, strconv.Itoa(int(userInfo.ID)))
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "登录成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 10003,
			"msg":  "登录失败",
		})
	}
}

func Logout(c *gin.Context) {
	session.ClearSession(c)
	c.Redirect(http.StatusFound, "/")
	return
}

func GetUserInfo(c *gin.Context) map[string]interface{} {
	return session.GetSessionUserInfo(c)
}
