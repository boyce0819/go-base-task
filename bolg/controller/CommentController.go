package controller

import (
	"github.com/gin-gonic/gin"
	"go-task1/bolg/config"
	"go-task1/bolg/errorss"
	"go-task1/bolg/model"
	"go-task1/bolg/results"
	"time"
)

//评论功能
//实现评论的创建功能，已认证的用户可以对文章发表评论。
//实现评论的读取功能，支持获取某篇文章的所有评论列表。
//错误处理与日志记录
//对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
//使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。

func CommentCreate(c *gin.Context) {
	var input model.CommentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}

	//组装数据
	comment := model.Comment{
		PostID:    input.PostID,
		Content:   input.Content,
		UserID:    input.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// 校验文章是否存在
	if err := config.Db.Select("id", input.PostID).First(&model.Post{}).Error; err != nil {
		results.Fail(c, errorss.ErrInvalidParams.Code, "post not exit")
	}
	if err := config.Db.Create(&comment).Error; err == nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	results.Success(c, nil, "success")
}

func CommentPostList(c *gin.Context) {
	var input model.CommentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	// 校验文章是否存在
	if err := config.Db.Select("id", input.PostID).First(&model.Post{}).Error; err != nil {
		results.Fail(c, errorss.ErrInvalidParams.Code, "post not exit")
	}
	var comments []model.Comment
	// 查询
	if err := config.Db.Select("post_id", input.PostID).First(&comments).Error; err != nil {
		results.Fail(c, errorss.ErrInvalidParams.Code, "post not exit")
	}
	results.Success(c, comments, "success")
}
