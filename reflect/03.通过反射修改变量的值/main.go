package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	// 整形修改值
	num := 18
	numRv := reflect.ValueOf(&num) //修改值必须传指针
	numRv.Elem().SetInt(20)
	fmt.Println("修改后num的值", num)

	// 结构体修改字段值
	stu := &Student{
		Id:   1,
		Name: "lili",
	}

	rv := reflect.ValueOf(stu)
	if rv.Kind() != reflect.Ptr {
		fmt.Println("修改变量值必须是指针类型")
	}
	// 获取指针指向的元素
	value := rv.Elem()
	// 获取Name字段对象
	nameObj := value.FieldByName("Name")
	// 修改Name字段的值
	if nameObj.Kind() == reflect.String {
		nameObj.SetString("zhangsan")
	}
	fmt.Printf("stu修改后：%+v", stu)
}
