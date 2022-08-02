package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// 产生问题的根本原因是golang的for循环会使用同一个变量来存储迭代过程中的临时变量，在将该变量传递给goroutine时，goroutine得到的是该变量的地址，
	// 又由于goroutine的启动与调度机制有关，可能for循环执行完后，goroutine才开始调度，所以导致多个goroutine访问的是同一个数据。
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("i：%v,addr：%p\n", i, &i)
		}()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
