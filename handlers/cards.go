package handlers

import (
	"groupie-pokemontcg/models"
	"html/template"
	"log"
	"net/http"
)

func add(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }

func CardsHandler(w http.ResponseWriter, r *http.Request) {

	//  On désactive tout ce qui vient de l’API
	// pageStr := r.URL.Query().Get("page")
	// sizeStr := r.URL.Query().Get("size")
	// typeFilter := r.URL.Query().Get("type")
	// hpFilter := r.URL.Query().Get("hp")
	// supertypeFilter := r.URL.Query().Get("supertype")

	//  On désactive la pagination dynamique
	// page := 1
	// size := 20

	//  On désactive l'appel API
	// cards, err := models.GetFilteredCards(page, size, typeFilter, hpFilter, supertypeFilter)
	// if err != nil {
	//     http.Error(w, "Erreur API", http.StatusInternalServerError)
	//     return
	// }

	// ✅ On force une data statique pour tester le front
	data := struct {
		Cards           []models.Card
		Page            int
		Size            int
		TypeFilter      string
		HPFilter        string
		SupertypeFilter string
	}{
		Cards: []models.Card{
			{
				ID:        "test-1",
				Name:      "Pikachu Debug",
				HP:        "60",
				Types:     []string{"Electric"},
				Supertype: "Pokémon",
				// Pas d'Images car ton modèle ne l'a pas
			},
		},
		Page:            1,
		Size:            20,
		TypeFilter:      "",
		HPFilter:        "",
		SupertypeFilter: "",
	}

	tmpl := template.Must(template.New("cards.html").Funcs(template.FuncMap{
		"add": add,
		"sub": sub,
	}).ParseFiles("templates/cards.html"))

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println("Erreur template:", err)
	}
}
