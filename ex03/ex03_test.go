package ex03

import (
	"reflect"
	"testing"
)

func TestNameAndAgeExist(t *testing.T) {
	typ := reflect.TypeOf(Person{})

	checkField(t, typ, "Name", "string", "")
	checkField(t, typ, "Age", "int", 0)
	checkIsAdultExists(t)
	checkMethod(t, typ, "IsValid", 0, 1)
}

func checkField(t *testing.T, typ reflect.Type, name string, typName string, typVal interface{}) {
	field, found := typ.FieldByName(name)
	if !found {
		t.Fatalf("Pas de champ '%s'", name)
	}
	if field.Type != reflect.TypeOf(typVal) {
		t.Fatalf("Le champ '%s' n'est pas de type %s", name, typName)
	}

}

func checkIsAdultExists(t *testing.T) {
	typ := reflect.TypeOf(Person{})

	method, found := typ.MethodByName("IsAdult")
	if !found {
		t.Fatalf("Pas de méthode 'IsAdult'")
	}

	ni := method.Type.NumIn()
	if ni != 1 {
		t.Fatalf("La méthode IsAdult ne doit prendre aucun paramètre")
	}
	no := method.Type.NumOut()
	if no != 1 {
		t.Fatalf("La méthode IsAdult doit renvoyer un seul résultat")
	}
	if method.Type.Out(0) != reflect.TypeOf(true) {
		t.Fatalf("La méthode IsAdult doit renvoyer un booléen (bool)")
	}

	p1 := Person{}
	vp1 := reflect.ValueOf(&p1)

	vp1.Elem().FieldByName("Age").Set(reflect.ValueOf(16))

	res := vp1.Elem().MethodByName("IsAdult").Call(nil)
	if res[0].Bool() {
		t.Fatalf("IsAdult a retourné true pour un age de 16")
	}

	vp1.Elem().FieldByName("Age").Set(reflect.ValueOf(18))

	res = vp1.Elem().MethodByName("IsAdult").Call(nil)
	if !res[0].Bool() {
		t.Fatalf("IsAdult a retourné false pour un age de 18")
	}

	vp1.Elem().FieldByName("Age").Set(reflect.ValueOf(42))

	res = vp1.Elem().MethodByName("IsAdult").Call(nil)
	if !res[0].Bool() {
		t.Fatalf("IsAdult a retourné false pour un age de 42")
	}

}

func checkMethod(t *testing.T, typ reflect.Type, name string, inCount, outCount int) {

	method, found := typ.MethodByName(name)
	if !found {
		t.Fatalf("Pas de méthode '%s'", name)
	}

	if method.Type.NumIn() != inCount+1 {
		t.Fatalf("La méthode %s doit prendre %d paramètre(s)", name, inCount)
	}

	if method.Type.NumOut() != outCount {
		t.Fatalf("La méthode %s doit renvoyer %d résultat(s)", name, outCount)
	}
	if method.Type.Out(0) != reflect.TypeOf(true) {
		t.Fatalf("La méthode IsAdult doit renvoyer un booléen (bool)")
	}

	p1 := Person{}
	vp1 := reflect.ValueOf(&p1)

	vp1.Elem().FieldByName("Age").Set(reflect.ValueOf(16))

	res := vp1.Elem().MethodByName("IsAdult").Call(nil)
	if res[0].Bool() {
		t.Fatalf("IsAdult a retourné true pour un age de 16")
	}

	vp1.Elem().FieldByName("Age").Set(reflect.ValueOf(18))

	res = vp1.Elem().MethodByName("IsAdult").Call(nil)
	if !res[0].Bool() {
		t.Fatalf("IsAdult a retourné false pour un age de 18")
	}

	vp1.Elem().FieldByName("Age").Set(reflect.ValueOf(42))

	res = vp1.Elem().MethodByName("IsAdult").Call(nil)
	if !res[0].Bool() {
		t.Fatalf("IsAdult a retourné false pour un age de 42")
	}

}
