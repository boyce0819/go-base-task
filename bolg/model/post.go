package model

import "time"

// Post posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
type Post struct {
	ID        int `gorm:"primary_key"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int `gorm:"foreignkey:UserID"`
}

type PostRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"  binding:"required"`
	Content string `json:"content"  binding:"required"`
}
