package ex03

import (
	"reflect"
	"testing"

	"github.com/DuchessFrance/hands-on-go/zutils"
)

func TestEx03(t *testing.T) {
	typ := reflect.TypeOf(Person{})

	zutils.CheckField(t, typ, "Name", "string", "")
	zutils.CheckField(t, typ, "Age", "int", 0)

	zutils.CheckMethod(t, typ, zutils.Method{Name: "IsAdult", OutTypes: zutils.Types(false)})
	p := zutils.NewStruct(&Person{})
	for _, cas := range []struct {
		age      int
		expected bool
	}{{16, false}, {18, true}, {42, true}, {0, false}} {
		if p.Set("Age", cas.age).Call("IsAdult")[0].Bool() != cas.expected {
			t.Fatalf("IsAdult a retourné %v pour un age de %d", !cas.expected, cas.age)
		}
	}

	zutils.CheckMethod(t, typ, zutils.Method{Name: "IsValid", OutTypes: zutils.Types(false)})
	for _, cas := range []struct {
		name     string
		age      int
		expected bool
		desc     string
	}{{"", 42, false, "nom vide"}, {"Pierre", -1, false, "age négatif"}, {"Pierre", 0, true, "age == 0"}, {"Pierre", 130, true, "age == 130"}, {"Pierre", 131, false, "age supérieur à 130"}} {
		p.Set("Name", cas.name).Set("Age", cas.age)
		if p.Call("IsValid")[0].Bool() != cas.expected {
			t.Fatalf("IsValid a retourné %v pour le cas: %v", !cas.expected, cas.desc)
		}
	}
}
