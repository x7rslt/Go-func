package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func Count() {
	keyword := `智能制造，数字化，智能化，信息化，自动化，云计算，云平台，物联网，数据可视化，信息物理系统，网络物理系统，大数据，感知技术，数据可视化，云制造，主动制造，智慧制造，智能企业，智能终端，智能技术，智能识别，机器人，工业4.0，工业互联网，互联网+，人机交互，数字智能，传感器，控制器，数据挖掘，数字技术，人工智能，人机协同，智能生产，智能生态系统，智能平台，传感网，超级计算，深度学习，群智开放，自主操控，人机互动，机器人，机器学习，自动分析，智能机器，算法，智能+，智能系统，智造云，云制造，算法分析`
	keywords := strings.Split(keyword, "，")

	pwd := "/Users/xiaoshuai/go/src/Project/goproject/excelChunk/output"

	//获取当前目录下的文件或目录名(包含路径)
	filepathNames, err := filepath.Glob(filepath.Join(pwd, "*"))
	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range filepathNames {
		fmt.Println("File:", filename)
		ok := strings.HasSuffix(filename, ".xlsx")
		if ok {
			f, err := excelize.OpenFile(filename)
			if err != nil {
				fmt.Println("filename:", filename, err)
				return
			}
			// Get all the rows in the Sheet1.
			rows, err := f.GetRows("Sheet1")
			for _, row := range rows {
				counters := make(map[string]int)
				content := strings.Split(row[5], " ")
				for _, i := range content {
					_, flag := Find(keywords, i)
					if flag {
						counters[i]++
					}
				}
				WritetoFile(counters, strings.Split(filename, ".")[0])
			}

		}

	}

}
func WritetoFile(m map[string]int, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, m)
	fmt.Println(filePath, "CountDone!")
	return w.Flush()
}

func main() {
	fmt.Println("Count")
	Count()
	fmt.Println("Complete.")
}
