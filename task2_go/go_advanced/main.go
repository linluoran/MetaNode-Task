package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var wg = sync.WaitGroup{}

func addNum(sNum *int) int {
	return *sNum + 10
}

func printOdd() {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
	}
}

func printEven() {
	defer wg.Done()
	for i := 2; i < 10; i += 2 {
		fmt.Println(i)
	}
}

type gaTask struct {
	name string
	job  func()
}

func schedule(tasks []gaTask) {
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(t gaTask) { // 使用参数传递避免闭包捕获问题
			defer wg.Done()
			startTime := time.Now()
			t.job()
			fmt.Printf("%s took %v\n", t.name, time.Since(startTime))
		}(task)
	}
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	res := r.width * r.height
	fmt.Println("方形面积为: ", res)
	return res
}

func (r Rectangle) Perimeter() float64 {
	res := 2 * (r.width + r.height)
	fmt.Println("方形周长为: ", res)
	return res
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	res := math.Pi * c.radius * c.radius
	fmt.Println("圆面积为: ", res)
	return res
}

func (c Circle) Perimeter() float64 {
	res := math.Pi * 2 * c.radius
	fmt.Println("圆周长为: ", res)
	return res
}

// 使用Shape接口的函数
func printShapeInfo(s Shape) {
	s.Area()
	s.Perimeter()
}

type person struct {
	name string
	age  int
}

type employee struct {
	person
	employeeID int
}

func (e employee) printInfo() {
	fmt.Println("员工的姓名: ", e.name)
	fmt.Println("员工的年纪: ", e.age)
	fmt.Printf("员工的ID: %09d\n", e.employeeID)
}

type numPro struct {
	num int
	mu  sync.Mutex
}

func (n *numPro) addNum(num int) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.num += num
}

func main() {

	// go进阶 1
	goAdvanced1 := 1
	fmt.Println(addNum(&goAdvanced1))

	// go进阶 2
	wg.Add(2)
	go printOdd()
	go printEven()
	wg.Wait()

	tasks := []gaTask{
		{"任务1", func() { time.Sleep(1 * time.Second) }},
		{"任务2", func() { time.Sleep(2 * time.Second) }},
		{"任务3", func() { time.Sleep(3 * time.Second) }},
		{"任务4", func() { time.Sleep(4 * time.Second) }},
	}
	schedule(tasks)
	wg.Wait()

	// go进阶 3 - 使用Shape接口
	rect := Rectangle{9.5, 11.5}
	printShapeInfo(rect)

	circle := Circle{3.5}
	printShapeInfo(circle)

	// go进阶 4
	lilei := employee{person{"lilei", 20}, 1502320}
	lilei.printInfo()

	// go进阶 5
	workCh := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(workCh)
		for i := 1; i < 11; i++ {
			workCh <- i
			fmt.Println("写入通道: ", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range workCh {
			fmt.Println("取出通道: ", res)
		}
	}()
	wg.Wait()

	// go进阶5
	cacheCh := make(chan int, 60)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(cacheCh)
		for i := 100; i < 200; i++ {
			cacheCh <- i
		}
		fmt.Println("数据写入完毕.")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range cacheCh {
			fmt.Println("从缓存通道中拿到: ", res)
		}
	}()
	wg.Wait()

	// go进阶 6
	wg.Add(10)
	countNum := &numPro{}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				countNum.addNum(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(countNum.num)

	// go进阶 7
	var atomicNum int64

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&atomicNum, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(atomicNum)
}
