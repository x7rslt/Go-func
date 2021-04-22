package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {
	t1 := time.Now()
	wb, err := xlsx.OpenFile("./exceltest/fenci.xlsx")
	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
	}
	fmt.Println("----")

	sheetName := "Sheet1"
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	fmt.Println("Max row in sheet:", sh.MaxRow)
	t2 := time.Since(t1)
	fmt.Println("Process cost time:", t2)
}
