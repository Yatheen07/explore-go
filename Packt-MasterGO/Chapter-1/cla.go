package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	var min float64 = 1000
	var max float64 = -1000
	arguments := os.Args
	if len(arguments) <=1{
		fmt.Println("Please provide two floats value")
		os.Exit(1)
	}
	for i:=1;i<len(arguments);i++{
		value,_ := strconv.ParseFloat(arguments[i],64)
		if value < min{
			min = value
		}
		if value > max{
			max = value
		}
	}
	fmt.Println(max,min)

}