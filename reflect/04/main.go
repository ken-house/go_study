package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

func (stu student) Hello(name string) string {
	return "hello " + name
}

func main() {
	stu := student{
		name: "lili",
		age:  18,
	}

	rv := reflect.ValueOf(stu)
	num := rv.NumMethod()
	fmt.Printf("获取到%d个方法", num)
	// 获取方法控制权
	methodValue := rv.MethodByName("Hello")
	// 拼凑参数
	args := []reflect.Value{reflect.ValueOf("reflect")}
	// 调用函数
	retList := methodValue.Call(args)
	fmt.Println(retList[0].String())
}
