package main

import (
	"io"
	"os"
)

func main(){
	input := ""
	arguments := os.Args
	if len(arguments) <= 1 {
		input = "Please include arguments"
	} else{
		input = arguments[1]
	}
	io.WriteString(os.Stdout,input)
	io.WriteString(os.Stdout,"\n")
}