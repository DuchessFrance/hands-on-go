package zutils

import (
	"reflect"
	"testing"
)

func CheckField(t *testing.T, typ reflect.Type, name string, typName string, typVal interface{}) {
	field, found := typ.FieldByName(name)
	if !found {
		t.Fatalf("Pas de champ '%s'", name)
	}
	if field.Type != reflect.TypeOf(typVal) {
		t.Fatalf("Le champ '%s' n'est pas de type %s", name, typName)
	}

}

type Method struct {
	Name     string
	InTypes  []reflect.Type
	OutTypes []reflect.Type
}

func CheckMethod(t *testing.T, typ reflect.Type, desc Method) {
	method, found := typ.MethodByName(desc.Name)
	if !found {
		t.Fatalf("Le type %v ne définit pas de méthode '%s'", typ, desc.Name)
	}

	if method.Type.NumIn() != len(desc.InTypes)+1 {
		t.Fatalf("La méthode %v.%s doit prendre %d paramètre(s)", typ, desc.Name, len(desc.InTypes))
	}

	if method.Type.NumOut() != len(desc.OutTypes) {
		t.Fatalf("La méthode %v.%s doit renvoyer %d résultat(s)", typ, desc.Name, len(desc.OutTypes))
	}

	for i, inType := range desc.InTypes {
		if method.Type.In(i+1) != inType {
			t.Fatalf("La méthode %v.%s doit prendre un type %v en paramètre #%d", typ, desc.Name, inType, i+1)
		}
	}

	for i, outType := range desc.OutTypes {
		if method.Type.Out(i) != outType {
			t.Fatalf("La méthode %v.%s doit renvoyer un type %v en résultat #%d", typ, desc.Name, outType, i+1)
		}
	}
}

func Types(values ...interface{}) []reflect.Type {
	res := make([]reflect.Type, len(values))
	for i, val := range values {
		res[i] = reflect.TypeOf(val)
	}
	return res
}
