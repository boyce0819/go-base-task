package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-task1/bolg/config"
	"go-task1/bolg/errorss"
	"go-task1/bolg/model"
	"go-task1/bolg/results"
	"go-task1/bolg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var input model.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	// 校验邮箱是否是规范邮箱， 插件正则 todo 自己手写怎么写， 注册到校验器中

	// 检查邮箱是否已存在
	var existingUser model.User
	if err := config.Db.Where("email = ?", input.Email).First(&existingUser).Error; err != nil {
		results.Fail(c, errorss.ErrSystem.Code, "邮箱已使用")
	}
	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	// 插入数据
	if err := config.Db.Create(&model.User{
		Email:    input.Email,
		Username: input.Username,
		Password: string(hashedPassword),
	}).Error; err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	results.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	var input model.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
	}
	// 查找用户
	var user model.User
	// 这里只写了邮箱登录验证， 查找用户，
	if err := config.Db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			results.Fail(c, errorss.ErrUserNotFound.Code, errorss.ErrUserNotFound.Error())
			return
		}
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		results.Fail(c, errorss.ErrInvalidParams.Code, "Invalid password")
		return
	}
	// 生成jwt
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		results.Fail(c, errorss.ErrSystem.Code, errorss.ErrSystem.Error())
		return
	}
	// 组装返回
	results.Success(c, token, "")
}
