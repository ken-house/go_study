package main

import (
	"fmt"
	"reflect"
)

type myInt int

type student struct {
	name string
	age  int
}

func main() {
	// 自定义类型
	var i myInt = 6
	rt := reflect.TypeOf(i)  // 读取变量（interface{}）的类型
	rv := reflect.ValueOf(i) // 读取变量（interface{}）的值
	fmt.Printf("rt:%+v,rv:%+v\n", rt, rv)
	fmt.Println("变量的原生基本类型：", rt.Kind(), rv.Kind())

	// 结构体类型
	stu := student{
		name: "lili",
		age:  18,
	}
	rt2 := reflect.TypeOf(stu)  // 读取变量（interface{}）的类型
	rv2 := reflect.ValueOf(stu) // 读取变量（interface{}）的值
	fmt.Printf("rt:%+v,rv:%+v\n", rt2, rv2)
	fmt.Println("变量的原生基本类型：", rt2.Kind(), rv2.Kind())

	// 判断变量类型
	if k := rv2.Kind(); k == reflect.Struct {
		fmt.Println("是我想要的类型-struct")
	}
}
