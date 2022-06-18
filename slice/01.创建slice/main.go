package main

import "fmt"

func main() {
	// 1、通过make创建
	sliceA := make([]int, 5, 10)
	fmt.Printf("通过make创建：长度为%d，容量为%d，slice：%+v\n", len(sliceA), cap(sliceA), sliceA)

	// 2、通过array创建
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceB := array[5:7]

	// slice容量为array[5]到数组结束
	fmt.Printf("通过数组创建：长度为%d，容量为%d，slice：%+v\n", len(sliceB), cap(sliceB), sliceB)

	// 数组和切片公用同一块内存，即&array[5]和&slice[0]地址相同
	fmt.Printf("array[5]的地址：%p\n", &array[5])
	fmt.Printf("sliceB[0]的地址：%p\n", &sliceB[0])

	// 3、通过slice创建与通过array创建效果一样
	sliceC := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceD := sliceC[5:7]
	fmt.Printf("通过slice创建：长度为%d，容量为%d，slice：%+v\n", len(sliceD), cap(sliceD), sliceD)
	fmt.Printf("sliceC[5]的地址：%p\n", &sliceC[5])
	fmt.Printf("sliceD[0]的地址：%p\n", &sliceD[0])
}
