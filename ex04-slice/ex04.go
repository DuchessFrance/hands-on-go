package ex04

// TODO: implémeter NewSlice de façon qu'elle renvoie un tableau avec un seul élément (élement passé en paramètre)
func NewSlice(element int) []int {
	return []int{element}
}

// TODO: implémeter LenSlice de façon qu'elle retourne le nombre d'éléments dans le tableau
func LenSlice(slice []int) int {
	return len(slice)
}

// TODO: implémeter AppendToSlice de façon qu'elle rajoute l'élément passé en paramètre à la fin du tableau
func AppendToSlice(slice []int, element int) []int {
	return append(slice, element)
}

// TODO: implémeter InSlice de façon qu'elle retourne true si le tableau slice contient l'élément element
func InSlice(slice []int, element int) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

// TODO: implémeter IndexInSlice de façon qu'elle retourne la position de l'élément si le tableau slice le contient, -1 sinon
func IndexInSlice(slice []int, element int) int {
	for idx, s := range slice {
		if s == element {
			return idx
		}
	}
	return -1
}
