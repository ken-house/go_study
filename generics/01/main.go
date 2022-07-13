package main

import "fmt"

// 注意看~int8
func add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	// 限制底层数据类型
	fmt.Println("MyInt 1 + 2 = ", add[int](1, 2))
	fmt.Println("MyInt 1 + 2 = ", add[float64](1.0, 2.2))
}
