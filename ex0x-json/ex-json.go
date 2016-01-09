package ex_json

import (
	"encoding/json"
	"io/ioutil"
)

//"path/filepath"
//"io/ioutil"
//"encoding/json"

type DuchessFrance struct {
	Organisation string
	Team         string
	Members      []People
}

type People struct {
	Name   string
	Status string
}

// TODO : Implémenter la fonction LoadJSON
func LoadJsonData() DuchessFrance {
	duchessFrance := DuchessFrance{}

	// TODO : Lire à l'aide de ioutil le fichier duchess.json
	dat, err := ioutil.ReadFile("duchess.json")
	if err != nil {
		panic(err)
	}

	// TODO : Unmarshal les données et remplir le struct duchessFrance

	json.Unmarshal(dat, &duchessFrance)

	return duchessFrance
}
