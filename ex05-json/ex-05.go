package ex_json

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

	// TODO 1 : Lire à l'aide de ioutil le fichier duchess.json

	// TODO 2 : Unmarshal les données et remplir le struct duchessFrance

	return nil
}
