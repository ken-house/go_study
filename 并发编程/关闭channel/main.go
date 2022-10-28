package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 当缓冲区写满时阻塞
			fmt.Printf("写入i=%v\n", i)
		}
		close(ch) // 明确知道已经没有数据会再发送到channel了，此时关闭channel
	}()

	// 不用知道要循环多少次，通过判断是否能从通道中拿到数据做判断
	for {
		if num, ok := <-ch; !ok {
			break
		} else {
			fmt.Printf("读取i=%v\n", num)
		}
	}
	fmt.Println("main goroutine 执行")
}
