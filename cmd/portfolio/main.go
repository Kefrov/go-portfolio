package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Kefrov/go-portfolio/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("public/static"))
	server.Handle("/static/", http.StripPrefix("/static/", fs))
	server.HandleFunc("/", handlers.HomeHandler)
	server.HandleFunc("/about", handlers.AboutHandler)
	server.HandleFunc("/projects", handlers.ProjectsHandler)
	server.HandleFunc("/contact", handlers.ContactHandler)

	log.Println("Server starting on port " + port)
	err := http.ListenAndServe(":"+port, server)
	if err != nil {
		log.Fatalf("Error while starting the server: %v", err)
	}
}
