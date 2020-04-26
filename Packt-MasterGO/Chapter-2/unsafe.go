package main

import (
	"fmt"
	"unsafe"
)

func printPointerDetails(pointerName string, pointer *int64){
	fmt.Printf("Address stored in %s is %p\n",pointerName,pointer)
	fmt.Printf("Value stored in %s is %d\n",pointerName,*pointer)
	fmt.Printf("Address of %s is %p\n",pointerName,&pointer)
	fmt.Println("=================================")
}

func printPointerDetails32(pointerName string, pointer *int32){
	fmt.Printf("Address stored in %s is %p\n",pointerName,pointer)
	fmt.Printf("Value stored in %s is %d\n",pointerName,*pointer)
	fmt.Printf("Address of %s is %p\n",pointerName,&pointer)
	fmt.Println("=================================")
}

func main(){
	var value int64 = 5;
	var pointer1 = &value
	printPointerDetails("Pointer 1",pointer1)	

	
	var pointer2 = (*int32) (unsafe.Pointer(pointer1))
	printPointerDetails32("Pointer 2",pointer2)
	*pointer1 = 61289736871236987 //Not a 32 bit integer
	printPointerDetails("Pointer 1",pointer1)
	printPointerDetails32("Pointer 2",pointer2)


}