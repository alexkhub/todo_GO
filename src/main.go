package main

import (
	"fmt"
	"net/http"
	"time"
	"src/handlers"
	
	
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/",  MainPageHandler)


	server := http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server start")
	server.ListenAndServe()
}