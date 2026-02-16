# TP-GROUPIE-TRACKER-API

- Compte rendu du projet — Groupie Tracker

Dans ce projet, j’ai dû créer un site web en Go qui utilise une API REST pour afficher des données. J’ai choisi l’API Pokémon TCG parce qu’elle est simple, bien documentée et qu’elle propose plusieurs endpoints différents, ce qui correspondait aux demandes du sujet.

Le site permet d’afficher une collection de cartes Pokémon, de voir les détails d’une carte, de faire une recherche, d’utiliser plusieurs filtres, d’avoir une pagination et de gérer une liste de favoris enregistrée dans un fichier JSON.

Pour la recherche, j’ai utilisé deux propriétés : le nom et le type.
Pour les filtres, j’en ai mis trois : type, rareté et set.

La pagination affiche les cartes par lots de 20.

Les favoris sont sauvegardés dans un fichier JSON pour qu’ils restent même après avoir relancé le serveur.

J’ai organisé mon projet avec des dossiers séparés : handlers, templates, models, static, etc. Ça m’a aidé à garder un code plus clair et plus facile à maintenir.

Pour l’API, j’ai surtout utilisé les endpoints /cards, /cards/{id}, /types et /sets. J’ai récupéré les données en JSON et je les ai traitées manuellement en Go.

Pour la gestion des erreurs, j’ai ajouté des vérifications simples pour éviter que le site plante si l’API ne répond pas ou si un fichier JSON est vide.

Au niveau de l’organisation, j’ai commencé par tester l’API, puis j’ai fait les pages HTML, ensuite j’ai mis en place les routes Go, et enfin j’ai ajouté les fonctionnalités une par une.
J’ai essayé de faire les choses dans l’ordre pour éviter de me perdre.

Ce projet m’a permis de mieux comprendre comment fonctionne une API, comment manipuler du JSON en Go, et comment structurer un petit site web sans framework. Même si c’était parfois compliqué, ça m’a vraiment aidé à progresser.

-Arborescence:

/data
    favorites.json
/handlers
    cards.go
    card_details.go
    favorites.go
    search.go
/models
    card.go
/static
    css/
    img/
/templates
    home.html
    cards.html
    details.html
    favorites.html
    search.html
    about.html
main.go