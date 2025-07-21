package main

import (
	"fmt"
	"math"
	"strconv"
	_ "strings"
	"sync"
	"time"
)

//指针
//题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
//考察点 ：指针的使用、值传递与引用传递的区别。

func updatePointValue(point *int) int {
	i := 10
	point = &i
	return *point
}

/*
DoubleSlice 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
切片与数组区别

主要区别总结
数组：长度固定，声明时确定
切片：长度可变，可以动态增长
数组：值传递（函数参数传递时会复制整个数组）
切片：引用传递（底层共享数组）
数组：在栈上分配（小数组）或堆上分配（大数组）
切片：在堆上分配，包含指向底层数组的指针、长度和容量
数组：[n]T（长度是类型的一部分）
切片：[]T（长度不是类型的一部分）
切片是基于数组的抽象，提供更灵活的操作
切片可以共享底层数组
使用建议
优先使用切片，因为它更灵活且功能更强大
只有在明确需要固定长度或性能敏感场景才使用数组
切片操作（如append）可能会触发底层数组的重新分配
*/

func DoubleSlice(slicePtr *[]int) {
	if slicePtr == nil {
		return
	}
	// x2
	slice := *slicePtr
	for i := 0; i < len(*slicePtr); i++ {
		slice[i] *= 2
	}
	// x2
	fmt.Println(slice)
	for i := range slice {
		slice[i] *= 2
	}
	fmt.Println(slice)
}

// Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func printOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 添加延迟使输出更清晰
	}
}

func printEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 添加延迟使输出更清晰
	}
}

//题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
//考察点 ：协程原理、并发任务调度。

// 找人 100 个人里找一个
func doGetNum(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for i := 1; i < 100; i++ {
		if i == num {
			fmt.Println("找到了：" + strconv.Itoa(num))
			end := time.Now()
			fmt.Println(end.Sub(start))
			return
		}
	}
	fmt.Println("抱歉：我未能找到了：" + strconv.Itoa(num))
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func getNumStu() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go doGetNum(i, &wg)
	}
	wg.Wait()
	fmt.Println("执行结束")
}

//面向对象
//题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
//实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//考察点 ：接口的定义与实现、面向对象编程风格。
/*
type <interface_name> interface {
   <method_name>(<method_params>) [<return_type>...]
   ...
}*/

/*
以下是实现Shape接口的完整Go代码，包含Rectangle和Circle两种形状的实现：

# Shape接口实现

关键点说明：
1 Shape接口声明了Area()和Perimeter()两个方法
3.Rectangle和Circle结构体分别实现了这两个方法
Go语言中实现接口是隐式的，只要实现了接口的所有方法即视为实现了该接口
printShapeInfo函数可以接受任何实现了Shape接口的类型
Shape类型的变量可以持有任何实现了该接口的值
*/
type Shape interface {
	Area() float64
	Perimeter() float64 // 周长
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

// Area 计算矩形面积  实现 Area
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius * 2
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	person     Person
}

func (e Employee) Print() {
	fmt.Printf("员工ID: %s\n姓名: %s\n年龄: %d\n",
		e.EmployeeID,
		e.person,
		e.person.Name,
		e.person.Age)
}

//Channel
//题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
//另一个协程从通道中接收这些整数并打印出来。
//考察点 ：通道的基本使用、协程间通信。

func printChannel() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 10; i++ {
			fmt.Printf("放入channel的值: %d\n", i)
			ch <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 从channel接收数据
		for value := range ch {
			fmt.Printf("接收到的值: %d\n", value)
		}
	}()
	wg.Wait()
}

//func basicChannel() {
//	fmt.Println("\n=== 基本channel示例 ===")
//	ch := make(chan int)
//
//	go func() {
//		fmt.Println("goroutine准备发送数据")
//		ch <- 42 // 发送数据到channel
//		fmt.Println("goroutine发送完成")
//	}()
//
//	fmt.Println("main等待接收数据")
//	value := <-ch // 从channel接收数据
//	fmt.Printf("接收到的值: %d\n", value)
//}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func bufferedChannel() {
	fmt.Println("\n=== 带缓冲channel示例 ===")
	ch := make(chan int, 10) // 缓冲大小为2
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 101; i++ {
			fmt.Printf("放入channel的值: %d\n", i)
			ch <- i
		}
	}()
	wg.Add(1)
	go func() {
		// 从channel接收数据
		defer wg.Done()
		for value := range ch {
			fmt.Printf("接收到的值: %d\n", value)
		}
	}()
	wg.Wait()
}

// SafeCounter 锁机制
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// SafeCounter 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Inc 自增
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	//atomic.AddInt64(&c.value, 1)
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func doInc() {
	var (
		wg      sync.WaitGroup
		counter SafeCounter
	)
	workerCount := 10
	iterations := 1000

	// 设置计数器初始值  有点可重入锁的味道
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func(workerId int) {
			// 每个计数器 减少
			defer wg.Done()
			for i := 1; i <= iterations; i++ {
				counter.Inc()
				// 模拟工作负载
				time.Sleep(5 * time.Microsecond)
			}
			fmt.Printf("Worker %d 完成\n", workerId)
		}(i)
	}
	wg.Wait()
	fmt.Printf("最终计数器值: %d (预期: %d)\n",
		counter.Value(), workerCount*iterations)
}

//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ：原子操作、并发数据安全。
