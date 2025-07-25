package orm

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// User 进阶gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
type User struct {
	ID         int `gorm:"primarykey"`
	Name       string
	Email      string
	Password   string
	PostsCount int       `gorm:"default:0"` // 新增文章统计字段
	CreatedAt  time.Time //创建时间
	UpdatedAt  time.Time // 修改时间
	Posts      []Post
}

type Post struct {
	ID        int `gorm:"primarykey"`
	UserID    int `gorm:"foreignkey:UserID"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment
}

type Comment struct {
	ID            int `gorm:"primarykey"`
	PostID        int `gorm:"foreignkey:PostID"`
	UserID        int `gorm:"foreignkey:UserID"`
	Data          string
	CommentCount  int `gorm:"default:0"`
	CommentStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	User          User `gorm:"foreignKey:UserID"`
	Post          Post `gorm:"foreignKey:PostID"`
}

// SelectUserAllPostAndComment 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
func SelectUserAllPostAndComment(db *gorm.DB, userId int) {
	// 建表
	InitTable(db)
	err := insertTestData(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	var posts []Post
	// 预加载用户的所有文章和每篇文章的评论
	preloadErr := db.Preload("Comments").
		Preload("Comments.User"). // 加载评论的用户信息
		Where("user_id = ?", userId).
		Find(&posts).Error
	if preloadErr != nil {
		fmt.Printf("查询用户文章失败: %d", preloadErr)
	}
	printPostsWithComments(posts)
}

// printPostsWithComments 打印文章和评论
func printPostsWithComments(posts []Post) {
	fmt.Printf("共查询到 %d 篇文章:\n", len(posts))
	for i, post := range posts {
		fmt.Printf("\n%d. 文章标题: %s\n", i+1, post.Title)
		fmt.Printf("   文章内容: %.50s...\n", post.Content)
		fmt.Printf("   评论数: %d\n", len(post.Comments))
		for j, comment := range post.Comments {
			fmt.Printf("   %d.%d 评论用户: %s\n", i+1, j+1, comment.User.Name)
			fmt.Printf("      评论内容: %.30s...\n", comment.Data)
		}
	}
}

// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
// Todo
func selectCommentBest(db *gorm.DB) {
	// "select * from posts where id = select posts_id from comments order by CommentCount desc limit 1"
	var post Post
	err := db.Raw(`SELECT a.* FROM posts a JOIN (
	      SELECT post_id , COUNT(*) as comment_count  GROUP BY post_id ORDER BY comment_count desc LIMIT 1 ) as  b
	     on a.id = b.post_id`).Scan(&post).Error
}

// AfterCreate 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(db *gorm.DB) error {
	// 更新相应用户的PostsCount字段
	result := db.Model(&User{}).Where("id = ?", p.UserID).Update("posts_count", gorm.Expr("posts_count + ?", 1))
	if result != nil {
		fmt.Print("Post钩子函数执行错误: d%", result)
		return result.Error
	}
	return nil
}

// BeforeDelete 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) BeforeDelete(db *gorm.DB) error {
	if c.CommentCount == 0 {
		result := db.Model(&Comment{}).Where("id = ?", c.ID).Update("comment_status", "无评论")
		if result != nil {
			fmt.Print("Comment钩子函数执行错误: d%", result)
			return result.Error
		}
	}
	return nil
}

//func dbInit() *gorm.DB {
//	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
//	// 打开数据库连接
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		//Logger: logger.Default.LogMode(logger.Info), // 打印所有SQL
//	})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	fmt.Println("Database connection established")
//	// 开启Debug模式（会打印所有SQL）
//	db = db.Debug()
//	// 获取底层 sql.DB 对象进行连接池配置
//	sqlDB, err := db.DB()
//	if err != nil {
//		panic("failed to get underlying sql.DB")
//	}
//	// 配置连接池
//	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
//	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
//	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间
//	return db
//}

func InitTable(db *gorm.DB) {
	err := autoMigrateBolgTable(db)
	if err != nil {
		return
	}
}

func autoMigrateBolgTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func insertTestData(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1. 创建测试用户
		users := []User{
			{
				Name:     "alice",
				Email:    "alice@example.com",
				Password: "alice123",
			},
			{
				Name:     "bob",
				Email:    "bob@example.com",
				Password: "bob123",
			},
			{
				Name:     "charlie",
				Email:    "charlie@example.com",
				Password: "charlie123",
			},
		}

		for i := range users {
			if err := tx.Create(&users[i]).Error; err != nil {
				return fmt.Errorf("创建用户失败: %w", err)
			}
		}

		// 2. 创建测试文章
		posts := []Post{
			{
				Title:   "Go语言入门指南",
				Content: "这是一篇关于Go语言基础知识的文章...",
				UserID:  users[0].ID,
			},
			{
				Title:   "GORM使用教程",
				Content: "本文将介绍如何使用GORM进行数据库操作...",
				UserID:  users[0].ID,
			},
			{
				Title:   "Web开发实战",
				Content: "使用Go语言构建Web应用的实践经验分享...",
				UserID:  users[1].ID,
			},
			{
				Title:   "微服务架构设计",
				Content: "如何设计一个高效的微服务架构系统...",
				UserID:  users[2].ID,
			},
		}

		for i := range posts {
			if err := tx.Create(&posts[i]).Error; err != nil {
				return fmt.Errorf("创建文章失败: %w", err)
			}
		}

		// 3. 创建测试评论
		comments := []Comment{
			{
				Data:   "非常好的入门教程！",
				UserID: users[1].ID,
				PostID: posts[0].ID,
			},
			{
				Data:   "期待更多关于GORM的高级用法",
				UserID: users[2].ID,
				PostID: posts[0].ID,
			},
			{
				Data:   "GORM确实很方便使用",
				UserID: users[1].ID,
				PostID: posts[1].ID,
			},
			{
				Data:   "Web开发实战对我帮助很大",
				UserID: users[0].ID,
				PostID: posts[2].ID,
			},
			{
				Data:   "微服务架构确实很复杂",
				UserID: users[0].ID,
				PostID: posts[3].ID,
			},
			{
				Data:   "有没有更详细的示例代码？",
				UserID: users[1].ID,
				PostID: posts[3].ID,
			},
		}

		for i := range comments {
			if err := tx.Create(&comments[i]).Error; err != nil {
				return fmt.Errorf("创建评论失败: %w", err)
			}
		}
		return nil
	})
}
