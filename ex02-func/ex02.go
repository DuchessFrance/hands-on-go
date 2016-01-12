package ex02

// TODO 1: déclarer une fonction "sum" qui accepte 2 entiers et qui retourne leur somme

func sum(a, b int) int {
	return a + b
}

// TODO 2: déclarer une fonction "predSucc" qui accepte un entier et qui renvoie son prédecesseur et son successeur

func predSucc(x int) (int, int) {
	return x - 1, x + 1
}
