package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	server := http.NewServeMux()

	server.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})

	log.Println("Server starting on port " + port)
	err := http.ListenAndServe(":" + port, server)
	if err != nil {
		log.Fatalf("Error while starting the server: %v", err)
	}
}
