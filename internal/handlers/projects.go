package handlers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Kefrov/go-portfolio/internal/data"
)

func mod(i, j int) bool {
	return i%j == 0
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("base").Funcs(template.FuncMap{
		"mod": mod,
	})

	html, err := tmpl.ParseFiles(
		"public/templates/base.html",
		"public/templates/projects.html",
		"public/templates/header.html",
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to parse templates: %v", err)
		return
	}

	err = html.ExecuteTemplate(w, "base", data.DataProjects)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to render templates: %v", err)
		return
	}
}
