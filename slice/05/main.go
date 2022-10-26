package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a == nil)
	fmt.Printf("自身地址：%p，指向堆地址：%p，值为：%v，长度：%v，容量：%v\n", &a, a, a, len(a), cap(a))
	a = append(a, 1)
	fmt.Println(a == nil)
	fmt.Printf("自身地址：%p，指向堆地址：%p，值为：%v，长度：%v，容量：%v\n", &a, a, a, len(a), cap(a))

	// 一次性append多个元素只会触发一次扩容，且扩容后的大小和长度有关（向上的偶数）。
	a = append(a, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Printf("自身地址：%p，指向堆地址：%p，值为：%v，长度：%v，容量：%v\n", &a, a, a, len(a), cap(a))
	//a = append(a, 1)
	//fmt.Printf("自身地址：%p，指向堆地址：%p，值为：%v，长度：%v，容量：%v\n", &a, a, a, len(a), cap(a))
	//a = append(a, 1)
	//fmt.Printf("自身地址：%p，指向堆地址：%p，值为：%v，长度：%v，容量：%v\n", &a, a, a, len(a), cap(a))
	//a = append(a, 1)
	//fmt.Printf("自身地址：%p，指向堆地址：%p，值为：%v，长度：%v，容量：%v\n", &a, a, a, len(a), cap(a))
}
