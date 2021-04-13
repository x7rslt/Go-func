package main

import (
	"fmt"
	"os"
	"strconv"
)

func Add(x, y int) {
	fmt.Println(x + y)
}

func Sqrt(x int) {

	fmt.Printf("%d Sqrt:%d \n", x, x*x)
}

func main() {
	input := os.Args[1:]
	if len(input) == 2 {
		x, _ := strconv.Atoi(input[0])
		y, _ := strconv.Atoi(input[1])
		Add(x, y)
	} else if len(input) == 1 {
		x, _ := strconv.Atoi(input[0])
		Sqrt(x)
	} else {
		fmt.Println("Error.Please input one or two number for calc.")
		os.Exit(0)
	}
}
