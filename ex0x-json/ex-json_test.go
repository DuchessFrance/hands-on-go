package ex_json

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExJson_01(t *testing.T) {

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
	require.Equal(t, expected, actual)
}
