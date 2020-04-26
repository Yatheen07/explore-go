package main

import(
	"log"
	"os"
)

func main(){
	outputFile,errorCode := os.OpenFile("logs/testLogFile.txt",os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if errorCode != nil{
		log.Fatalf("Error Opening File: %v", errorCode)
	}
	defer outputFile.Close()

	log.SetOutput(outputFile)
	log.Println("Test log entry")
}