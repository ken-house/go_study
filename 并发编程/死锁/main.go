package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ch = make(chan int)
var rwMux sync.RWMutex

func read(i int) {
	for {
		rwMux.RLock()
		num := <-ch
		fmt.Printf("第%v次读取随机数%v\n", i, num)
		rwMux.RUnlock()
	}
}

func write(i int) {
	for {
		num := rand.Intn(1000)
		rwMux.Lock()
		ch <- num
		fmt.Printf("第%v次产生随机数%v\n", i, num)
		rwMux.Unlock()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 5; i++ {
		go read(i)
	}

	for i := 1; i <= 5; i++ {
		go write(i)
	}

	for {
	}
}
