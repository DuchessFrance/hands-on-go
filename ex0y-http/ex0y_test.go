package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"encoding/json"

	"reflect"
)

func TestHello(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Fatalf("Erreur de connexion au serveur sur localhost:8080: %v", err)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erreur de lecture sur la réponse de /hello: %v", err)
	}

	if string(contents) != "Hello World" {
		t.Fatalf("/hello a renvoyé %q plutôt que %q", contents, "Hello World")
	}
}

func TestDuchesses(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/duchesses")
	if err != nil {
		t.Fatalf("Erreur de connexion au serveur sur localhost:8080: %v", err)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Erreur de lecture sur la réponse de /duchesses: %v", err)
	}

	var actual []Duchess
	err = json.Unmarshal(contents, &actual)
	if err != nil {
		t.Fatalf("Problème de parsing du json de /duchesses: %v", err)
	}

	if !reflect.DeepEqual(duchesses, actual) {
		t.Fatalf("Le json renvoyé par /duchesses est incorrect")
	}
}
