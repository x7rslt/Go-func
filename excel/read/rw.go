package main

import (
	"bufio"
	"fmt"
	"os"
)

func WriteMaptoFile(m map[string]int, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k, v := range m {
		lineStr := fmt.Sprintf("%s:%d", k, v)
		fmt.Fprintln(w, lineStr)
	}
	return w.Flush()
}
func main() {
	path := "counter.txt"
	content := map[string]int{"信息化": 1, "智能化": 1, "机器人": 1}
	WriteMaptoFile(content, path)
}
