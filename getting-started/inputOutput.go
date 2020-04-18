package main

import (
	"fmt"
	"strings"
)

func sample() string {
	return "Sample"
}

func main() {
	var input float64
	fmt.Println("Enter input:")
	fmt.Scanf("%f", &input) //same old style C -_-

	output := input * 2
	fmt.Println(output)

	name := "Yatheen"
	for _,value := range name {
		fmt.Printf("%c\n",value); //This way you have to format and print. 
	}

	split := strings.Split(name,"") //This is split the string into a rune slice
	for index,value := range split{
		temp := value
		fmt.Println(index,temp)
	}
}
