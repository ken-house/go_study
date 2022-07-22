package main

import (
	"log"
	"net/http"
	"os"
)

// 设置日志记录器
func setupLogger() {
	logFileLocation, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

// 简单发送HTTP请求
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
		return
	}
	log.Printf("Status Code for %s : %s", url, resp.Status)
	resp.Body.Close()
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("success")
}

func main() {
	//defer fmt.Println("panic退出前处理")
	//log.Println("println日志")
	//log.Panic("panic日志")
	//log.Fatal("程序退出日志")

	log.Println("println日志")

	//setupLogger()
	//simpleHttpGet("www.baidu.com")
	//simpleHttpGet("http://www.baidu.com")
}
