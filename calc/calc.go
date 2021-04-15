package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func Add(x, y int) {
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}

func Sqrt(x int) {
	fmt.Printf("%d的平方：%d\n", x, x*x)
}

var add string
var sqrt int

func main() {
	flag.StringVar(&add, "add", "", "-add num1,num2")
	flag.IntVar(&sqrt, "sqrt", 7, "-sqrt num")
	flag.Parse()
	if add != "" {
		value := strings.Split(add, ",")
		x, _ := strconv.Atoi(value[0])
		y, _ := strconv.Atoi(value[1])
		Add(x, y)
	} else {
		Sqrt(sqrt)
	}
}
