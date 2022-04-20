package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	AvatarId  string    `json:"avatar_id"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05" json:"created_at"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05" json:"updated_at"`
	DeletedAt time.Time `time_format:"2006-01-02 15:04:05" json:"deleted_at"`
}

// GetUserByField 根据ID或者用户名获取用户信息
func GetUserByField(field, value string) User {
	var u User
	if field == "id" || field == "username" {
		DB.Where(field+" = ?", value).First(&u)
	}
	return u
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(avatarId string, u User) User {
	u.AvatarId = avatarId
	DB.Save(u)
	return u
}

// AddUser 添加用户
func AddUser(value interface{}) User {
	var u User
	u.Username = value.(map[string]interface{})["username"].(string)
	u.Password = value.(map[string]interface{})["password"].(string)
	u.AvatarId = value.(map[string]interface{})["avatar_id"].(string)
	DB.Create(&u)
	return u
}
