package main
import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func main(){
	arguments := os.Args
	var result float64
	var numericCounter float64 = 0
	if len(arguments) == 1{
		fmt.Println("No arguments provided")
		os.Exit(1)
	} else{
		for i := 1;i<len(arguments);i++{
			value, err := strconv.ParseFloat(arguments[i],64)
			if err  == nil{
				result += value
				numericCounter++
			}
		}
	}
	if numericCounter == 0{
		fmt.Println("No numeric arguments found")
	} else{
		fmt.Printf("Total sum: %v\n",result)
		fmt.Printf("Average: %v\n", math.Abs(result/numericCounter) )
	}
}