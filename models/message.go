package models

import "time"

type Message struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	RoomId    int64     `json:"room_id"`
	ToUserId  int64     `json:"to_user_id"`
	Content   string    `json:"content"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05" json:"created_at"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05" json:"updated_at"`
	DeletedAt time.Time `time_format:"2006-01-02 15:04:05" json:"deleted_at"`
}
