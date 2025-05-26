package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 是一个线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex // 互斥锁，用于保护 count 的并发访问
	count int        // 实际的计数值
}

// Increment 安全地增加计数器（线程安全）
func (c *SafeCounter) Increment() {
	c.mu.Lock()         // 加锁，防止其他 goroutine 同时修改
	defer c.mu.Unlock() // 函数返回时自动解锁
	c.count++           // 增加计数
}

// GetCount 安全地获取当前计数值（线程安全）
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()         // 加锁，防止其他 goroutine 同时读取
	defer c.mu.Unlock() // 函数返回时自动解锁
	return c.count      // 返回当前计数值
}

// UnsafeCounter 是一个非线程安全的计数器
type UnsafeCounter struct {
	count int // 计数值，没有锁保护
}

// Increment 非安全地增加计数器（并发不安全）
func (c *UnsafeCounter) Increment() {
	c.count += 1 // 直接修改 count，可能导致数据竞争（race condition）
}

// GetCount 非安全地获取当前计数值（并发不安全）
func (c *UnsafeCounter) GetCount() int {
	return c.count // 直接返回 count，可能读取到中间状态
}

func main() {
	counter := UnsafeCounter{} // 创建一个非线程安全的计数器

	// 启动 1000 个 goroutine 并发增加计数器
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment() // 每个 goroutine 增加 100 次
			}
		}()
	}

	// 等待 1 秒，确保所有 goroutine 完成（实际生产环境应该用 sync.WaitGroup）
	time.Sleep(time.Second)

	// 输出最终计数（由于并发不安全，结果可能小于 100000）
	fmt.Printf("Final count: %d (expected: 100000)\n", counter.GetCount())

	// 对比 SafeCounter 的正确行为
	safeCounter := SafeCounter{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				safeCounter.Increment() // 线程安全地增加计数
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("Safe final count: %d (correct)\n", safeCounter.GetCount())
}
