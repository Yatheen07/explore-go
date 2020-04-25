package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var input *os.File
	input = os.Stdin
	defer input.Close()

	fmt.Printf("> ")
	scanner := bufio.NewScanner(input)
	for scanner.Scan(){		
		currentInput := scanner.Text()
		if currentInput == "stop"{
			break
		} else{
			fmt.Printf("> %s\n",currentInput)
		}
		fmt.Printf(">")
	}
}