package main

import "net/http"

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

	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//TODO 1: ecrire Hello World dans la response
}

// TODO 3: implémenter un nouveau handler pour GET /duchesses
