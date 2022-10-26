package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type studentList []Student

// studentList实现sort.Interface接口
func (s studentList) Len() int {
	return len(s)
}

func (s studentList) Less(i, j int) bool {
	return s[i].Age > s[j].Age
}

func (s studentList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var list studentList
	stu1 := Student{
		Name: "张三",
		Age:  10,
	}
	// 这里可以定义多个...
	stu10 := Student{
		Name: "李十",
		Age:  20,
	}
	list = append(list, stu1, stu10)

	// 直接将list作为参数传入
	sort.Sort(list)
	fmt.Println(list)
}
