package main

import (
	"fmt"
	"runtime"
	"time"
	"os"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC) //This line prints the number of garbage collectors currently active now.
	fmt.Println("-----")
}

func main(){
	var memoryObject runtime.MemStats
	printStats(memoryObject)

	for i:=0;i<5;i++{
		sliceVariable := make([]byte,50000000)
		if sliceVariable == nil{
			fmt.Println("Operation Failed!");
			os.Exit(1)
		}
	}
	printStats(memoryObject)

	for i:=0;i<5;i++{
		sliceVariable := make([]byte,10000000)
		if sliceVariable == nil{
			fmt.Println("Operation Failed!");
			os.Exit(1)
		}
		time.Sleep(5 * time.Second)
	}
	printStats(memoryObject)
}