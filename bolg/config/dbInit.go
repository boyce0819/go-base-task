package config

import (
	"fmt"
	"go-task1/bolg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

func DbInit() {
	dsn := "root:root@tcp(127.0.0.1:3306)/bolg?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info), // 打印所有SQL
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection established")
	// 开启Debug模式（会打印所有SQL）
	db = db.Debug()
	// 获取底层 sql.DB 对象进行连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get underlying sql.DB")
	}
	// 配置连接池
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	// 初始化表
	initTables(db)
}

func initTables(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
