package ex04

import (
	"reflect"
	"testing"
)

func TestEx04(t *testing.T) {
	actual := NewSlice(42)
	if len(actual) != 1 {
		t.Fatalf("le slice retourné par NewSlice n'a pas la bonne taille (%d au lieu de 1)", len(actual))
	}

	if actual[0] != 42 {
		t.Fatalf("le slice retourné par NewSlice ne contient pas le paramètre element")
	}

	if LenSlice(nil) != 0 {
		t.Fatalf("LenSlice ne retourne pas 0 pour un tableau vide")
	}

	if LenSlice(make([]int, 36521)) != 36521 {
		t.Fatalf("LenSlice ne retourne pas la bonne taille")
	}

	appended := AppendToSlice([]int{1, 2, 3}, 42)
	if len(appended) != 4 {
		t.Fatalf("le slice retourné par AppendToSlice n'a pas la bonne taille")
	}

	if !reflect.DeepEqual(appended, []int{1, 2, 3, 42}) {
		t.Fatalf("le slice retourné ne contient pas l'élément passé en paramètre")
	}

	if InSlice(nil, 5) {
		t.Fatalf("InSlice retourne true alors que le tableau passé est vide")
	}

	if !InSlice([]int{1, 2, 5}, 5) {
		t.Fatalf("InSlice retourne false alors que le tableau contient bien l'élément")
	}

	if InSlice([]int{1, 2, 3}, 5) {
		t.Fatalf("InSlice retourne true alors que le tableau ne contient pas l'élément")
	}

	if IndexInSlice(nil, 5) != -1 {
		t.Fatalf("IndexInSlice ne retourne pas -1 alors que le tableau passé est vide")
	}

	if IndexInSlice([]int{1, 2, 5}, 5) != 2 {
		t.Fatalf("IndexInSlice retourne pas la bonne position")
	}

	if IndexInSlice([]int{1, 2, 3}, 5) != -1 {
		t.Fatalf("IndexInSlice ne retourne pas -1 alors que le tableau ne contient pas l'élément")
	}
}
