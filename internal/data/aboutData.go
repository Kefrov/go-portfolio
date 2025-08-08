package data

import (
	"html/template"

	"github.com/Kefrov/go-portfolio/internal/models"
)

var DataAbout = models.AboutData{
	ActivePage: 1,
	Paragraphs: []template.HTML{
		`I'm currently studying at <span class="highlight">ISAMM</span>, slowly building my foundations and improving my intuition. So far, I’ve always been the <span class="highlight">top student</span> in my bachelor's program. I've also made the greatest friends there!`,
		`I enjoy <span class="highlight">mathematics</span> and <span class="highlight">problem solving</span>. It's never fun to just learn, I like puzzles. I've reached the <span class="highlight">Expert</span> rank on Codeforces. However, recently I've been trying to apply those skills and build more.`,
		`In my spare time, I like reading math books, listening to music, gaming, and messing around with friends.`,
	},
}
