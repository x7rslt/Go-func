package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//确认是否包含关键字
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//字频统计
func Count(filelist []string) {
	keywords := strings.Split(ConfigData.Keywords, "，")
	for i, filename := range filelist {
		fmt.Printf("Count NO.%d:【%s】 >>>>>> ", i, strings.Split(filename, "/")[len(strings.Split(filename, "/"))-1])
		f, err := excelize.OpenFile(filename)
		if err != nil {
			fmt.Println("Error open filename:", filename, err)
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
			WritetoFile(counters, strings.Split(filename, ".xlsx")[0]+".txt")
		}
		fmt.Println("CountDone!")

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
	return w.Flush()
}

func main() {
	fmt.Println("------Start Count------")
	filelist := Getfilelist()
	Count(filelist)
	fmt.Println("------Complete!------")
}
