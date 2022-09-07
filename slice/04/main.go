package main

import "fmt"

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func test(data []User) {
	fmt.Printf("%p,%+v\n", data, data)
	person := User{
		Id:   1,
		Name: "zhangsan",
	}
	data[0] = person
	//*data = append(*data, person)
	fmt.Printf("%p,%+v\n", data, data[0])
}

func main() {
	data := make([]User, 1, 100)
	fmt.Printf("%p,%+v\n", data, data)
	test(data)
	fmt.Printf("%p,%+v\n", data, data[0])
}
