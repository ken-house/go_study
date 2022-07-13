package main

import "fmt"

func main() {
	sliceA := []int{1, 2}
	fmt.Printf("扩容前容量为：%d，地址为：%p\n", cap(sliceA), &sliceA[0])

	// 通过append添加元素，由于内存不足，进行扩容即分配一个更大的内容空间，因此地址会发生变化；
	// 扩容策略为当前容量小于1024，则扩容为原来的2倍，若扩容后容量小于所需容量，则容量为所需容量；
	// 这里当前容量为2，新增3个元素，理论上应该为5个容量，但实际内存分配还需要考虑内存对齐，因此分配了6个容量；
	sliceA = append(sliceA, 3, 4, 5)
	fmt.Printf("扩容后容量为：%d，地址为：%p\n", cap(sliceA), &sliceA[0])

	// 对于一个未分配内存空间的slice，通过append可以开辟一个内存空间，因此这里不会报错；
	// 第二次append时容量已扩容到6，当第三次添加元素时，不需要扩容，因此第二次和第三次的地址是同一个
	var sliceB []int
	sliceB = append(sliceB, 1, 2, 3)
	fmt.Printf("扩容前容量为：%d，地址为：%p\n", cap(sliceB), &sliceB[0])
	sliceB = append(sliceB, 4, 5)
	fmt.Printf("扩容后容量为：%d，地址为：%p\n", cap(sliceB), &sliceB[0])
	sliceB = append(sliceB, 6)
	fmt.Printf("扩容后容量为：%d，地址为：%p\n", cap(sliceB), &sliceB[0])
}
