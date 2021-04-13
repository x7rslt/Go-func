package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func command() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Command:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	input = strings.Replace(input, "\n", "", -1)
	inputlist := strings.Split(input, " ")
	cmd := exec.Command(inputlist[0], inputlist[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		return
	}
	fmt.Println("Command executioned:")
	fmt.Println("-----------------------------")
	fmt.Println(string(out))
}

func main() {
	for {
		command()
	}
}
