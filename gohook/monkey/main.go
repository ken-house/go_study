package main

import (
	"bou.ke/monkey"
	"fmt"
)

func originFunc() {
	str := "Hi, origin func"
	fmt.Println(str)
}

func monkeyFunc() {
	str := "Hi, monkey func"
	fmt.Println(str)
}

func main() {
	originFunc()
	monkey.Patch(originFunc, monkeyFunc)
	originFunc()
}
