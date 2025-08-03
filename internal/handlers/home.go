package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles(
		"public/templates/base.html",
		"public/templates/home.html",
		"public/templates/header.html",
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to parse templates: %v", err)
		return
	}

	err = html.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to render templates: %v", err)
		return
	}
}
