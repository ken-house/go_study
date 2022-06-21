package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	apiUrl := "http://127.0.0.1:8081/hello"
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Printf("http.NewRequest err:%v", err)
		return
	}

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("client.Do err:%v", err)
				return
			}

			if resp.Body == nil {
				fmt.Printf("resp.Body == nil")
				return
			}

			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("ioutil.ReadAll:%v", err)
				return
			}

			fmt.Println(string(b))
		}()
	}
	wg.Wait()
}
