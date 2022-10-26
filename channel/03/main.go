package main

import (
	"fmt"
)

var ch = make(chan int, 50)
var exitCh = make(chan bool, 1)

func writer() {
	for i := 1; i <= 50; i++ {
		ch <- i
	}

	// 关闭channel
	close(ch)
}

func reader() {
	for v := range ch {
		fmt.Println(v)
	}
	// 读取完成后往exitCh中写入数据
	exitCh <- true
}

func main() {
	go writer()
	go reader()

	// 循环等待，直到从exitCh管道中拿到数据则退出程序
	for {
		if _, ok := <-exitCh; ok {
			break
		}
	}
}
