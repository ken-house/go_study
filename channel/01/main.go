package main

import "fmt"

func main() {
	myChan := make(chan int, 3)
	// 向channel发送一个值
	myChan <- 1
	myChan <- 1
	myChan <- 1
	myChan <- 1
	fmt.Printf("channel的长度：%v，容量：%v\n", len(myChan), cap(myChan))
	// 从channel中接收一个值
	num := <-myChan
	fmt.Println(num)
	num = <-myChan
	fmt.Println(num)
}
