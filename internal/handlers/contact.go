package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/Kefrov/go-portfolio/internal/data"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	// handles form submission
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error Parsing Form", http.StatusBadRequest)
			log.Printf("Failed to parse form: %v", err)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		subject := fmt.Sprintf("Portfolio Contact from %s", name)
		body := fmt.Sprintf("From: %s\nEmail: %s\n\nMessage:\n%s", name, email, message)

		// create a mailto URL to open the user's email client
		mailtoURL := fmt.Sprintf("mailto:%s?subject=%s&body=%s",
			data.DataContact.Email,
			url.PathEscape(subject),
			url.PathEscape(body),
		)

		http.Redirect(w, r, mailtoURL, http.StatusSeeOther)
		return
	}

	html, err := template.ParseFiles(
		"public/templates/base.html",
		"public/templates/contact.html",
		"public/templates/header.html",
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to parse templates: %v", err)
		return
	}

	err = html.ExecuteTemplate(w, "base", data.DataContact)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Failed to render templates: %v", err)
		return
	}
}
