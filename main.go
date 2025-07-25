package main

import (
	"fmt"
	"go-task1/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)

	//for i := 1; i <= 5; i++ {
	////TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
	//// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
	//fmt.Println("i =", 100/i)
	//}

	// 第一题找数组的唯一元素
	//a := []int{5, 5, 2, 2, 1}
	//var a []int // 数组没有初始化直接传到方法里，方法里没做处理不报错？？ 写Java的有点接受不了
	//fmt.Println(getOnlyOnceNum(a))

	// 回文数
	//fmt.Println(isPalindrome(12321))

	// 最长公共前缀
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	//task2
	// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值
	//var p3 *int
	//p := 2
	//p3 = &p
	//fmt.Println(updatePointValue(p3))

	// DoubleSlice 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	// 测试用例
	//DoubleSlice(&[]int{1, 2, 3})

	// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	// 考察点 ： go 关键字的使用、协程的并发执行。 go 关键字启用携程
	// sync.WaitGroup 猜测就是一个线程安全的计数器，通过add 及done 方法修改
	//var wg sync.WaitGroup
	//wg.Add(2) // 等待两个协程完成
	//// 为何要用指针作为参数，不用值呢， 指针地址安全，值多携程修改 不安全
	//go printOddNumbers(&wg)
	//go printOddNumbers(&wg)
	//wg.Wait() // 等待所有协程完成
	//fmt.Println("所有数字打印完成")

	//  找人 100 个人里找一个
	//getNumStu()

	//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
	// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	//e := Employee{
	//	EmployeeID: 1,
	//	person: Person{
	//		Name: "Boyce",
	//		Age:  18,
	//	},
	//}
	//e.Print()

	// 普通channel
	//printChannel()
	// 缓冲channel
	//bufferedChannel()

	//锁机制
	//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	//考察点 ： sync.Mutex 的使用、并发数据安全。
	//doInc()
	//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	//考察点 ：原子操作、并发数据安全。
	// 	//atomic.AddInt64(&c.value, 1) 修改inc 方法，使用工具包atomic AddInt64  返回的value 值也用atomic 包一下

	/*gorm----------------------------------------------------------------------*/
	// MySQL 连接配置
	//orm.RunDb(dbInit())
	// 测试 RunStudent
	//orm.RunStudent(dbInit())
	//orm.InitData(dbInit())
	//orm.EmployeesSelect(dbInit())
	//orm.BookSelect(dbInit())
	/*---------------------------------bolg-------------------------------------*/
	//orm.InitTable(dbInit())
	//orm.InitTable(dbInit())
	orm.SelectUserAllPostAndComment(dbInit(), 1)
}

func dbInit() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
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
	return db
}
