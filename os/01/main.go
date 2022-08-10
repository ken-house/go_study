package main

import (
	"os"
	"path/filepath"
)

func main() {
	// 判断目录是否存在，不存在则创建
	filePath := "./a/b/c/1.txt"
	dirPath := filepath.Dir(filePath)
	_, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				return
			}
			_, err = os.Create(filePath)
			return
		}
	}
	//f, err := os.Open(".")
	//if err != nil {
	//	return
	//}
	//for {
	//	dirList, err := f.ReadDir(-1)
	//	if err != nil {
	//		return
	//	}
	//	for _, v := range dirList {
	//		fmt.Printf("v.IsDir(): %v，v.Name(): %v\n", v.IsDir(), v.Name())
	//	}
	//}

}
