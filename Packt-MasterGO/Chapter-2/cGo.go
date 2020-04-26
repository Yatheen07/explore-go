package main

//#include <stdio.h>
//void cMethodName() {
// printf("Inside C code\n");
//}
import "C"
import "fmt"

func main(){
	fmt.Println("From go!")
	C.cMethodName()
	fmt.Println("From go2!")
}