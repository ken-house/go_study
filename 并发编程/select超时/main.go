package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select { // 读数据
			case num := <-ch:
				fmt.Println("读取到", num)
			case nowTimer := <-time.After(time.Second * 2): // 设置定时器
				fmt.Println(nowTimer)
				quit <- true
			}
		}
	}()

	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("main over")
}
