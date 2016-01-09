package ex04

// TODO 1: implémeter NewSlice de façon qu'elle renvoie un tableau avec un seul élément (élement passé en paramètre)
func NewSlice(element int) []int {
	return []int{element}
}

// TODO 2: implémeter LenSlice de façon qu'elle retourne le nombre d'éléments dans le tableau
func LenSlice(slice []int) int {
	return len(slice)
}

// TODO 3: implémeter AppendToSlice de façon qu'elle rajoute l'élément passé en paramètre à la fin du tableau
func AppendToSlice(slice []int, element int) []int {
	return append(slice, element)
}

// TODO 4: implémeter InSlice de façon qu'elle retourne true si le tableau slice contient le paramètre 'element'
func InSlice(slice []int, element int) bool {
	for _, s := range slice {
		if s == element {
			return true
		}
	}
	return false
}

// TODO 5: implémeter IndexInSlice de façon qu'elle retourne la position de l'élément si le tableau slice le contient, -1 sinon
func IndexInSlice(slice []int, element int) int {
	for idx, s := range slice {
		if s == element {
			return idx
		}
	}
	return -1
}
