package orm

import (
	"fmt"
	"gorm.io/gorm"
)

// Student SQL语句练习
// 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
type Student struct {
	ID    uint   // Standard field for the primary key
	Name  string // A regular string field
	Age   uint8  // An unsigned 8-bit integer
	Grade string // A regular string field
}
type Result struct {
	Name string
	Age  int
}

//func RunStudent(db *gorm.DB) {
//	var err = autoMigrate(db)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
//	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
//	//student := Student{Name: "boyce", Age: 16, Grade: "七年级"}
//	//students := []*Student{
//	//	{Name: "shy", Age: 16, Grade: "七年级"},
//	//	{Name: "lisi", Age: 23, Grade: "七年级"},
//	//}
//	//db.Create(students)
//	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
//	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
//	//var stus []*Student
//	//db.Where("age > ?", "18").Find(&stus)
//	//for i, student := range stus {
//	//	fmt.Printf("%d: ID=%d, Name=%s\n", i+1, student.ID, student.Name)
//	//	// 可以修改学生信息
//	//	//student.Age += 1
//	//	// 然后可以 db.Save(student) 保存修改
//	//}
//
//	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
//	// 全量更新 db.save
//	// 查出 姓名为 "张三" 的学生 的对象  r
//	//var student Student
//	//db.Where("name = ?", "张三").Find(&student)
//	//student.Grade = "四年级"
//	//db.Save(&student)
//	// 根据条件更新
//	// 根据条件更新
//	// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
//	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;
//	//db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
//
//	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
//	db.Where("age < ?", 15).Delete(&Student{})
//}

// Account 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
// 并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
type Account struct {
	ID      uint
	Name    string
	Balance float64
}

type Transactions struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func InitData(db *gorm.DB) {
	// 建表
	err := autoMigrate(db)
	if err != nil {
		return
	}
	// 创建账户
	accounts := []*Account{
		{Name: "A", Balance: 100},
		{Name: "B", Balance: 300},
	}
	for _, account := range accounts {
		err = db.Create(account).Error
	}

	if err := transferMoney(db, "A", "B", 100); err != nil {
		fmt.Println(err)
	}
}

func transferMoney(db *gorm.DB, from, to string, amount float64) error {
	//开启事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	fromAccount := Account{}
	// 需要先检查账户 A 的余额是否足够
	//tx.First(&Account{}, from).Scan(&toAccount)
	if err := tx.Where("name = ?", from).First(&fromAccount).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return fmt.Errorf("账户 %s 不存在", from)
	}
	if fromAccount.Balance < amount {
		fmt.Printf("%s账户金额不够\n", from)
		tx.Rollback()
		return fmt.Errorf("%s账户金额不够", from)
	}
	toAccount := Account{}
	// 校验b 账户存在？
	if err := tx.Where("name = ?", to).First(&toAccount).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return fmt.Errorf("账户 %s 不存在", to)
	}
	// 转账
	// a 扣 钱 ， b加钱 ， 记录表记录 A 向账户 B 转账 100 元的操作
	if err := tx.Model(&Account{}).Where("id = ?", fromAccount.ID).Update("amount ", fromAccount.Balance-amount).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return fmt.Errorf("更新%s账户出错", from)
	}

	if err := tx.Model(&Account{}).Where("id = ?", toAccount.ID).Update("amount ", toAccount.Balance+amount).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return fmt.Errorf("更新%s账户出错", to)
	}
	if err := tx.Create(&Transactions{
		FromAccountId: fromAccount.ID,
		ToAccountId:   toAccount.ID,
		Amount:        amount}).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	// 提交事务
	return tx.Commit().Error
}

// Employees Sqlx入门
// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
type Employees struct {
	ID         uint
	Name       string
	Age        uint8
	Department string
	Salary     float64
}

type EmployeesResult struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Department string
	Salary     float64
}

func EmployeesSelect(db *gorm.DB) {
	err := autoMigrate(db)
	if err != nil {
		return
	}

	// 创建账户
	employeeData := []*Employees{
		{Name: "A", Department: "技术部", Salary: 100},
		{Name: "B", Department: "技术部", Salary: 200},
		{Name: "C", Department: "技术部", Salary: 58999},
		{Name: "D", Department: "技术部", Salary: 58999},
		{Name: "E", Department: "人事部", Salary: 200},
		{Name: "F", Department: "人事部", Salary: 200},
	}
	for _, emp := range employeeData {
		err = db.Create(emp).Error
	}

	var employees []*Employees
	db.Where("department", "技术部").Scan(&employees)
	for employee := range employees {
		fmt.Println(employee)
	}
	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	var employee Employees
	db.Limit(1).Order("salary DESC").Find(&employee)
	fmt.Println(employee)
	fmt.Println(employee.Salary)
}

// Book 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func BookSelect(db *gorm.DB) {
	err := autoMigrate(db)
	if err != nil {
		return
	}
	// 创建账户
	books := []*Book{
		{Title: "科学", Author: "boyce", Price: 100},
		{Title: "体育", Author: "alice", Price: 200},
		{Title: "艺术", Author: "jerry", Price: 187},
		{Title: "漫画", Author: "九把刀", Price: 22},
	}
	for _, book := range books {
		err = db.Create(book).Error
	}

	var bookResultList []*Book
	db.Where("price > ?", 50).Find(&bookResultList)
	for _, bookItem := range bookResultList {
		fmt.Println(bookItem)
	}
}

//进阶gorm
//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//要求 ：
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
//编写Go代码，使用Gorm创建这些模型对应的数据库表。
//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。
//题目3：钩子函数
//继续使用博客系统的模型。
//要求 ：
//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Student{}, &Transactions{}, &Account{}, &Employees{}, &Book{})
}
