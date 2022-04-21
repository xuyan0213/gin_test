package session

import (
	"gin/config"
	"gin/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(config.Viper.GetString("app.cookie_key")))
	return sessions.Sessions("gin-test", store)
}

// SaveUserSession 注册和登录时保存用户session
func SaveUserSession(c *gin.Context, info interface{}) {
	session := sessions.Default(c)
	session.Set("u_id", info)
	err := session.Save()
	if err != nil {
		log.Fatal("登录信息保存失败:", err)
	}
}

// GetSessionUserInfo 通过session获取用户ID并查询出用户基本信息
func GetSessionUserInfo(c *gin.Context) map[string]interface{} {
	session := sessions.Default(c)
	uid := session.Get("u_id")
	data := make(map[string]interface{})
	if uid != nil {
		user := models.GetUserByField("id", uid.(string))
		data["u_id"] = user.ID
		data["username"] = user.Username
		data["avatar_id"] = user.AvatarId
	}
	return data
}

// ClearSession 退出时清除session
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

// CheckSession 检查Session是否存在并返回结果
func CheckSession(c *gin.Context) string {
	session := sessions.Default(c)
	uid := session.Get("u_id")
	if uid == nil {
		return ""
	}
	return uid.(string)
}

// SessionMiddleware session中间件
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := CheckSession(c)
		if uid == "" {
			c.Redirect(http.StatusFound, "/")
			return
		}
		uidInt, err := strconv.Atoi(uid)
		if err != nil {
			return
		}
		if uidInt <= 0 {
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.Set("uid", uid)
		c.Next()
		return
	}
}
