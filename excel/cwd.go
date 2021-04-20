package main

import (
	"fmt"
	"strings"

	//"io/ioutil"
	"log"
	"path/filepath"
)

func GetCwd() {
	pwd := "/Users/xiaoshuai/go/src/Project/goproject/excelChunk/output"

	//获取当前目录下的文件或目录名(包含路径)
	filepathNames, err := filepath.Glob(filepath.Join(pwd, "*"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filepathNames)
	var filelist []string
	for _, i := range filepathNames {
		ok := strings.HasSuffix(i, ".xlsx")
		if ok {
			filelist = append(filelist, i)
		}
	}
	fmt.Println(filelist)
}
func main() {
	GetCwd()
}
