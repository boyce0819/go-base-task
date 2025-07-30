package model

import "time"

//comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。

type Comment struct {
	ID        int `gorm:"primary_key"`
	UserID    int `gorm:"ForeignKey:UserID"`
	PostID    int `gorm:"ForeignKey:PostID"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentRequest struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	PostID    int    `json:"postId"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
