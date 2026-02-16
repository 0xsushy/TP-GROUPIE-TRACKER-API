package handlers

import (
	"encoding/json"
	"groupie-pokemontcg/models"
	"html/template"
	"io"
	"net/http"
	"os"
)

type Favorites struct {
	Favorites []string `json:"favorites"`
}

func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	file, err := os.Open("data/favorites.json")
	if err != nil {
		http.Error(w, "Impossible de lire les favoris", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var fav Favorites
	json.Unmarshal(bytes, &fav)

	// éviter les doublons
	for _, f := range fav.Favorites {
		if f == id {
			http.Redirect(w, r, "/favorites", http.StatusSeeOther)
			return
		}
	}

	fav.Favorites = append(fav.Favorites, id)

	newData, _ := json.MarshalIndent(fav, "", "  ")
	os.WriteFile("data/favorites.json", newData, 0644)

	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}
func FavoritesPageHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data/favorites.json")
	if err != nil {
		http.Error(w, "Impossible de lire les favoris", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var fav Favorites
	json.Unmarshal(bytes, &fav)

	var cards []models.Card
	for _, id := range fav.Favorites {
		card, err := models.GetCardByID(id)
		if err == nil {
			cards = append(cards, card)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/favorites.html"))
	tmpl.Execute(w, cards)
}
func RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	file, err := os.Open("data/favorites.json")
	if err != nil {
		http.Error(w, "Impossible de lire les favoris", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bytes, _ := io.ReadAll(file)

	var fav Favorites
	json.Unmarshal(bytes, &fav)

	var newFav []string
	for _, f := range fav.Favorites {
		if f != id {
			newFav = append(newFav, f)
		}
	}

	fav.Favorites = newFav

	newData, _ := json.MarshalIndent(fav, "", "  ")
	os.WriteFile("data/favorites.json", newData, 0644)

	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}
