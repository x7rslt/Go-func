//读取文本，提取URL连接，输出适合Awvs批量导入扫描的格式

package urlprase_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestFileRead(t *testing.T) {
	f, err := os.Open("./content.html")
	if err != nil {
		fmt.Println("Read file fail", err)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	var content []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		content = append(content, line)
		if err != nil {
			if err != nil {
				if err == io.EOF {
					fmt.Println("File read ok!")
					break
				} else {
					fmt.Println("Read file error!", err)
					return
				}
			}
		}
	}
	//fmt.Println(content)

	r, _ := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}`)
	f2, err := os.OpenFile("./url_list.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		t.Log(err)
	}
	var url string
	for _, u := range content {
		url = r.FindString(u)
		fmt.Println(url)
		_, err = f2.WriteString(url + "\n")
		if err != nil {
			t.Log(err)
		}
	}
}

func incr(p int) int {
	p++
	return p
}
func pointerincr(p *int) int {
	*p++
	return *p
}
func TestPoint(t *testing.T) {
	n := 1
	incr(n)
	fmt.Println(incr(n))
	//var p int
	pointerincr(&n)
	p := &n
	fmt.Printf("%T\n", *p)
	fmt.Println(pointerincr(&n))
}
