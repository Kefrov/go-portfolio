package data

import "github.com/Kefrov/go-portfolio/internal/models"

var DataHome = models.HomeData{
	ActivePage:  0,
	Name:        "Fares",
	Surname:     "Manai",
	Subtitle:    "I'm a CS student.",
	Description: "Thanks for visiting. This portfolio is my first web project. I am looking to become a well-rounded software engineer.",
	Socials: models.Socials{
		GitHub:    "https://github.com/Kefrov",
		LinkedIn:  "https://linkedin.com/in/kefrov",
		Instagram: "https://instagram.com/kefrov",
	},
}
