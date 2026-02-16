package main

import (
	"groupie-pokemontcg/handlers"
	"html/template"
	"log"
	"net/http"
)

func add(x, y int) int { return x + y }
func sub(x, y int) int { return x - y }

func main() {
	tmplFuncs := template.FuncMap{
		"add": add,
		"sub": sub,
	}

	templates := template.Must(template.New("").Funcs(tmplFuncs).ParseGlob("templates/*.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", nil)
	})

	http.HandleFunc("/cards", handlers.CardsHandler)
	http.HandleFunc("/card", handlers.CardDetailsHandler)
	http.HandleFunc("/add-favorite", handlers.AddFavoriteHandler)
	http.HandleFunc("/favorites", handlers.FavoritesPageHandler)
	http.HandleFunc("/remove-favorite", handlers.RemoveFavoriteHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
