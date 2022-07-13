package main

import "fmt"

func main() {
	sliceA := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceB := []int{6, 6, 6, 6, 6}

	// 将sliceA拷贝到sliceB上，只会拷贝前五个元素
	copy(sliceB, sliceA)

	fmt.Printf("sliceA:%v,cap:%d,len:%d\n", sliceA, cap(sliceA), len(sliceA))
	fmt.Printf("sliceB:%v,cap:%d,len:%d\n", sliceB, cap(sliceB), len(sliceB))

	// 将sliceB拷贝到sliceA上，将会把sliceB的5个元素全部拷贝过去
	sliceC := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceD := []int{6, 6, 6, 6, 6}
	copy(sliceC, sliceD)

	fmt.Printf("sliceC:%v,cap:%d,len:%d\n", sliceC, cap(sliceC), len(sliceC))
	fmt.Printf("sliceD:%v,cap:%d,len:%d\n", sliceD, cap(sliceD), len(sliceD))
}
