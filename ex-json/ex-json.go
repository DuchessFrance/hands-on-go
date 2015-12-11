package ex_json
import (
	//"path/filepath"
	//"io/ioutil"
	//"encoding/json"
)

type DuchessFrance struct {
	Organisation string
	Team string
	Members []People
}

type People struct {
	Name string
	Status string
}

// TODO : Implémenter la méthode LoadJSON
func LoadJsonData() DuchessFrance {
	duchessFrance := DuchessFrance{}

	// TODO : Lire à l'aide de filepath et ioutil le fichier duchess.json
	//absPath, _ := filepath.Abs("duchess.json")
	//dat, err := ioutil.ReadFile(absPath)
	//if (err != nil) {
	//	panic(err)
	//}

	// TODO : Unmarshal les données et remplir le struct duchessFrance


	//json.Unmarshal(dat, &duchessFrance)

	return duchessFrance
}
