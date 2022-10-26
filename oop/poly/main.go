package main

import "fmt"

type usb interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name + "手机开始工作")
}

func (p Phone) stop() {
	fmt.Println(p.Name + "手机停止工作")
}

type Carma struct {
	Name string
}

func (p Carma) start() {
	fmt.Println(p.Name + "相机开始工作")
}

func (p Carma) stop() {
	fmt.Println(p.Name + "相机停止工作")
}

type Computer struct {
}

func (c Computer) Work(usb usb) {
	usb.start()
	usb.stop()
}

func main() {
	// 定义一个手机
	phone := Phone{Name: "苹果"}
	// 定义一个相机
	carma := Carma{Name: "佳能"}

	computer := Computer{}
	computer.Work(phone)
	computer.Work(carma)
}
