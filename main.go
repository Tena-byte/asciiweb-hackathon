package main

import (
	"main/asciiweb"
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", asciiweb.EndPoints)
	mux.HandleFunc("/hello", asciiweb.GetName)
	mux.HandleFunc("/count", asciiweb.Count)
	mux.HandleFunc("/calculate", asciiweb.Calculate)
	mux.HandleFunc("/agent", asciiweb.Agent)


	fmt.Println("server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println("Server not running")
	}
}
