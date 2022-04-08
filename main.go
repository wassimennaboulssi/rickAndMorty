package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/wassimennaboulssi/RickAndMortyApi/rickAndMortyApi"
)

func main() {
	http.HandleFunc("/", CharactersListHandler)
	fs := http.FileServer(http.Dir("rickAndMortyApi/templates/css.css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.ListenAndServe(":8000", nil)
}

func CharactersListHandler(w http.ResponseWriter, r *http.Request) {
	results, _ := rickAndMortyApi.GetAllCharacters()
	CollectLocations(results)
	tmpl, _ := template.New("").ParseFiles("rickAndMortyApi/templates/characters.html", "rickAndMortyApi/templates/base.html")

	tmpl.ExecuteTemplate(w, "base", results)
}

func CollectLocations(results []rickAndMortyApi.Character) {

	resu := make(chan []rickAndMortyApi.LocationData)
	for _, character := range results {
		go func(character rickAndMortyApi.Character) {
			location, err := rickAndMortyApi.GetLocationData(character)
			if err != nil {
				log.Fatalf("error getting location data: %v", err)
			}

			resu <- location

		}(character)
	}

}
