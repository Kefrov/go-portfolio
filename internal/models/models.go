package models

import "html/template"

// Home
type Socials struct {
	GitHub    string
	LinkedIn  string
	Instagram string
}

type HomeData struct {
	ActivePage  int
	Name        string
	Surname     string
	Subtitle    string
	Description string
	Socials     Socials
}

// About
type AboutData struct {
	ActivePage int
	Paragraphs []template.HTML
}

// Projects
type Project struct {
	Title       string
	Url         string
	Description string
	ImagePath   string // aspect ratio should be 1:1
}

type ProjectsData struct {
	ActivePage int
	Projects   []Project
}

// Contact
type ContactData struct {
	ActivePage int
	Email      string
	Message    string
}
