package main

import(
	"fmt"
	"net/http"
	"log"
	"time"
)

func index_handler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Hello Worlds, %s!",request.URL.Path[1:])
}

func about_handler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "This is a second page, %s!",request.URL.Path[1:])
}

func main(){

	//This is a hanlder function, basically a url to function mapping
	http.HandleFunc("/",index_handler)
	//http.HandleFunc("/about/",about_handler)

	fmt.Println("Intiliasing the server...")
	
	//This makes the program listen in the specified port.
	//Arguments: Address and Handler,handler by default is nil
	//http.ListenAndServe(":8060",about_handler)

	//Custom Server Configuration
	server := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}