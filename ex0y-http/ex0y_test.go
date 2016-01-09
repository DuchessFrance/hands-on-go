package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/require"
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

	require.Equal(t, "Hello World", string(contents))
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
	require.NoError(t, err, "Problème de parsing du json de /duchesses: %v", err)

	require.Equal(t, duchesses, actual)
}
