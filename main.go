package main

import (
	"fmt"
	"net/http"
	"weeklytest/routes"
)

func main() {

	mux := http.NewServeMux()
	// r := &http.Request{}
	routes.Routes(mux)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Printf("server running at %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("err")
	}
}
