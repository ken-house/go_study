package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个监听的全局条件变量
var done = false

func read(name string, cond *sync.Cond) {
	// 先加锁
	cond.L.Lock()
	// 若done不为true，则继续等待
	if !done {
		cond.Wait()
	}
	// 若done为true，则解锁执行代码
	fmt.Println("start：" + name)
	cond.L.Unlock()
}

func write(name string, cond *sync.Cond) {
	fmt.Println("start：" + name)
	time.Sleep(time.Second)
	// 变量设置为true
	done = true
	// 广播唤醒所有等待goroutine
	cond.Broadcast()
}

func main() {
	// 创建一个cond基于互斥锁的条件变量
	cond := sync.NewCond(&sync.Mutex{})

	// 启动三个goroutine，每个goroutine都在等待变量done为true
	go read("read1", cond)
	go read("read2", cond)
	go read("read3", cond)

	// 执行一个操作
	write("write", cond)
	// 防止主进程退出
	time.Sleep(5 * time.Second)
}
