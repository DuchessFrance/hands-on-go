package ex03

/*
TODO: déclarer un struct Person:
*/

type Person struct {
	// TODO: 1. Ajouter un champ Name de type string
	Name string
	// TODO: 2. Ajouter un champ Age de type int
	Age int
}

// TODO: 3. Déclarer une méthode IsAdult sur Person qui renvoie true quand l'age est supérieur ou égal à 18, false sinon
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// TODO: 4. Déclarer une méthode IsAdult sur Person qui renvoie true quand l'age est supérieur ou égal à 18, false sinon
func (p Person) IsValid() bool {
	if len(p.Name) == 0 {
		return false
	}
	if p.Age < 0 {
		return false
	}
	if p.Age > 130 {
		return false
	}
	return true
}
