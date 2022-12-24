package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printPrompt() {
	fmt.Printf("DogDB > ")
}

func main() {
	inputStream := bufio.NewReader(os.Stdin)
	for true {
		printPrompt()
		command, _ := inputStream.ReadString('\n')

		// win 按下回车带有\r\n
		command = command[:len(command)-2]
		if strings.Compare(command, "exit") == 0 {
			break
		}
		fmt.Println(command)
	}
}
