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

	rt := reflect.TypeOf(stu)
	rv := reflect.ValueOf(stu)

	// 获取结构体数据信息
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("第%d个字段：字段名：%s，类型为：%s，标签为：%s，值为：%v\n", i+1, rt.Field(i).Name, rt.Field(i).Type, rt.Field(i).Tag.Get("json"), rv.Field(i))
	}

	// 获取结构体方法信息
	for i := 0; i < rt.NumMethod(); i++ {
		fmt.Printf("第%d个方法：方法名：%s，类型：%v\n", i, rt.Method(i).Name, rt.Method(i).Type)
	}
}
