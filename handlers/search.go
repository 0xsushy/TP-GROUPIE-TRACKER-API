package handlers

import (
	"groupie-pokemontcg/models"
	"html/template"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	var cards []models.Card
	if query != "" {
		cards, _ = models.SearchCards(query)
	}

	data := struct {
		Cards []models.Card
		Query string
	}{
		Cards: cards,
		Query: query,
	}

	tmpl := template.Must(template.ParseFiles("templates/search.html"))
	tmpl.Execute(w, data)
}
