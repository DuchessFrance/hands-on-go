package zutils

import "reflect"

type Wrapper struct {
	val reflect.Value
}

func NewStruct(val interface{}) *Wrapper {
	return &Wrapper{reflect.ValueOf(val)}
}

func (w *Wrapper) Set(field string, value interface{}) *Wrapper {
	w.val.Elem().FieldByName(field).Set(reflect.ValueOf(value))
	return w
}

func (w *Wrapper) Call(method string, args ...interface{}) []reflect.Value {
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	return w.val.Elem().MethodByName(method).Call(in)
}
