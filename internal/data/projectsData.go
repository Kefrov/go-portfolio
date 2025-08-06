package data

import "github.com/Kefrov/go-portfolio/internal/models"

var DataProjects = models.ProjectsData{
	ActivePage: 2,
	Projects: []models.Project{
		{
			Title:       "Blast",
			Url:         "https://github.com/Kefrov/Blast",
			Description: "This is my remake of the popular mobile game \"Block Blast\". It features a clean and simple UI, some animations, and a scoring system similar to the original. I used Python's Pygame library for the game logic and C++ to encrypt the high score.",
			ImagePath:   "/static/images/projects/blast.png",
		},
		{
			Title:       "Graphing Tool",
			Url:         "https://github.com/Kefrov/Graphing-tool",
			Description: "A math graphing tool I created in high school, inspired by GeoGebra. It was a fun project that took only two days to build, and it was the first time I realized how much I enjoy problem solving.",
			ImagePath:   "/static/images/projects/graph.png",
		},
		{
			Title:       "SOE",
			Url:         "https://github.com/Kefrov/SOE",
			Description: "This is a small Pygame app that demonstrates how the Sieve of Eratosthenes algorithm works. As the training manager for the problem-solving club at ISAMM, I used this tool to help my friends better understand the algorithm.",
			ImagePath:   "/static/images/projects/soe.png",
		},
		{
			Title:       "PIXSAM",
			Url:         "https://github.com/Kefrov/PIXSAM",
			Description: "An Android app I developed with a friend for a university project using Android Studio, primarily coded in Java. It's a simple pixel art app.",
			ImagePath:   "/static/images/projects/pixsam.png",
		},
	},
}
