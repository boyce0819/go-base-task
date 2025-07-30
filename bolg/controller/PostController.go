package controller

import (
	"github.com/gin-gonic/gin"
	"go-task1/bolg/config"
	"go-task1/bolg/errorss"
	"go-task1/bolg/model"
	"go-task1/bolg/results"
	"time"
)

func SeyHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "say hello",
	})
}

func SayGoodbye(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "say Goodbye",
	})
}

//文章管理功能
//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
//**实现文章的读取功能**，支持获取所有文章列表和单个文章的详细信息。
//实现文章的更新功能，只有文章的作者才能更新自己的文章。
//实现文章的删除功能，只有文章的作者才能删除自己的文章。

func CreatePost(c *gin.Context) {
	var input model.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	//校验用户是否存在 jwt 已校验
	// 创建数据
	post := model.Post{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := config.Db.Create(&post).Error; err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	results.Success(c, nil, "post created successfully")
}

func UpdatePost(c *gin.Context) {
	var input model.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	//校验用户是否存在 jwt 已校验
	// 创建数据
	post := model.Post{
		ID:        input.ID,
		Title:     input.Title,
		Content:   input.Content,
		UpdatedAt: time.Now(),
	}
	if err := config.Db.Save(&post).Error; err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	results.Success(c, nil, "post update successfully")
}

func DeletePost(c *gin.Context) {
	var input model.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	if err := config.Db.Delete("id", input.ID).Error; err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	results.Success(c, nil, "post delete successfully")
}

func GetPostDetail(c *gin.Context) {
	var input model.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	var post model.Post
	if err := config.Db.Select("id", input.ID).Scan(&post).Error; err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	results.Success(c, post, "")
}

func ListPost(c *gin.Context) {
	var posts []model.Post
	if err := config.Db.Find(&posts).Error; err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	results.Success(c, posts, "")
}
