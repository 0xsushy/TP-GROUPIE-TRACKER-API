package handlers

import (
    "html/template"
    "log"
    "net/http"

    "groupie-pokemontcg/models"
)

func CardDetailsHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID manquant", http.StatusBadRequest)
        return
    }

    card, err := models.GetCardByID(id)
    if err != nil {
        http.Error(w, "Carte introuvable", http.StatusNotFound)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/details.html"))
    err = tmpl.Execute(w, card)
    if err != nil {
        log.Println("Erreur template:", err)
    }
}