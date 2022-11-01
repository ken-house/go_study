package main

import (
	"fmt"
	"time"
)

// ----------定义一个task对象（任务）-------------------------
type task struct {
}

func NewTask() *task {
	return &task{}
}

func (t *task) execute() string {
	return fmt.Sprintf("打印当前时间为：%v", time.Now().Unix())
}

// -----------定义goroutine对象，从jobCh中读取任务执行---------
type work struct {
}

func NewWork() *work {
	return &work{}
}

func (w work) running(jobCh chan *task, i int) {
	for t := range jobCh {
		fmt.Printf("第%v个goroutine，%v\n", i, t.execute())
	}
}

func main() {
	// 定义一个接收task对象的channel
	taskCh := make(chan *task)
	// 定义一个提供给goroutine的channel
	jobCh := make(chan *task)
	// 指定N个goroutine
	workNum := 3

	// 将任务循环写入到taskCh通道中
	go func() {
		for {
			t := NewTask()
			taskCh <- t
		}
	}()

	// 启动指定个数的goroutine
	for i := 0; i < workNum; i++ {
		go func(i int) {
			w := NewWork()
			w.running(jobCh, i)
		}(i)
	}

	// 将taskCh通道里的任务写入到jobCh通道
	for t := range taskCh {
		jobCh <- t
	}
}
