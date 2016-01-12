package ex_json

import (
	"testing"

	"reflect"
)

func TestExJson01(t *testing.T) {

	actual := LoadJsonData()

	expected := DuchessFrance{
		Organisation: "Duchess France",
		Team:         "Go",
		Members: []People{
			{
				Name:   "Jawher Moussa",
				Status: "freelance",
			},
			{
				Name:   "Katia Aresti",
				Status: "freelance",
			},
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Le struct renvoyé par LoadJsonData est incorrect")
	}
}
