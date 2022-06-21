package main

import (
	"fmt"
	"time"
)

var ch = make(chan struct{})

func read(name string, ch chan struct{}) {
	// 从channel读取数据，阻塞在这里，直到读取到数据
	<-ch
	fmt.Println("start：" + name)
}

func write(name string, ch chan struct{}) {
	fmt.Println("start：" + name)
	time.Sleep(time.Second)
	// 关闭channel，这样其他goroutine就可以从channel读取数据
	close(ch)
}

func main() {
	// 启动三个goroutine，每个goroutine都在等待变量done为true
	go read("read1", ch)
	go read("read2", ch)
	go read("read3", ch)

	// 执行一个操作
	write("write", ch)
	// 防止主进程退出
	time.Sleep(5 * time.Second)
}
