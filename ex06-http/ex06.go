package main

import (
	"encoding/json"
	"net/http"
)

type Duchess struct {
	Name   string
	Status string
}

var duchesses = []Duchess{
	{Name: "JoJo", Status: "freelance"},
	{Name: "Karesti", Status: "freelance"},
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	// TODO 2: Ajouter route /duchesses qui renvoie une liste de Duchess en json
	http.HandleFunc("/duchesses", duchessesHandler)

	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//TODO 1: ecrire Hello World dans la response
	w.Write([]byte("Hello World"))
}

// TODO 2: implénter le handler
func duchessesHandler(w http.ResponseWriter, r *http.Request) {
	//TODO 1: ecrire Hello world dans la response
	contents, _ := json.Marshal(duchesses)
	w.Write(contents)
}
