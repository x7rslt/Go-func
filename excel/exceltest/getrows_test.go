package getrows_test

import (
	"fmt"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func TestGetRows(t *testing.T) {
	f, err := excelize.OpenFile("keyword.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		fmt.Println(row[5])
		fmt.Println()
	}
}
