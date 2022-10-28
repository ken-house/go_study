package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() { // 写数据
		for {
			select { // 读数据
			case num := <-ch:
				fmt.Println("读取到", num)
			case <-quit:
				return
			}
		}
	}()

	for i := 0; i < 2; i++ {
		ch <- i
	}
	quit <- true
}
