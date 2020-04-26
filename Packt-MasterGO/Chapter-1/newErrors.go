package main

import (
	"fmt"
	"errors"
)

func canDivideNumbers(a,b int) error{
	if a == 0 || b == 0{
		err := errors.New("Cant divide by 0 Exception")
		return err
	}else{
		return nil
	} 
}

func main(){
	err := canDivideNumbers(5,10)
	if err != nil{
		fmt.Println(err.Error())
	} else{
		fmt.Println("Yes, compatible!")
	}
	err = canDivideNumbers(5,0)
	if err != nil{
		fmt.Println(err.Error())
	} else{
		fmt.Println("Yes, compatible!")
	}
}