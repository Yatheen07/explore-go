package main

import (
	"fmt"
	"strings"
	"time"
	"log"
)

func computeExecutedTime(startTime time.Time, function string){
	elapsed := time.Since(startTime)
	log.Printf("Execution Time for %s : %s",function,elapsed)
}

func compareStrings(str1 string,str2 string){
	defer computeExecutedTime(time.Now(),"Own implentation")
	var result int
	for i:=0;i<len(str1) && i<len(str2);i++{
		if(str1[i] != str2[i]){
			result = int(str1[i] - str2[i])
			break
		}
	}
	time.Sleep(2 * time.Second)
	fmt.Println(result)
}

func main(){
	str1 := "Yatheen"
	str2 := "YatheenSampleString"

	//Built- Comparison Operators
	fmt.Println((str1 == str2),(str1 < str2),(str1 > str2))

	//string.compare method
	fmt.Println(strings.Compare(str1,str2))

	//Own implementation
	compareStrings(str1,str2)
}
