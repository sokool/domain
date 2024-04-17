package domain

import (
	"fmt"
	"reflect"
	"strconv"
)

type Traverser[T any] func(T) (bool, error)

func Travers[T any](v any, t Traverser[T]) error {
	rv := reflect.ValueOf(v)
	return travers(rv, rv.Type().Name(), t)
}

func travers[T any](rv reflect.Value, n string, fn Traverser[T]) error {
	if rv.IsZero() {
		return nil
	}
	var t, rt = [0]T{}, rv.Type()

	fmt.Println(n)
	if rt.AssignableTo(reflect.TypeOf(t).Elem()) {
		ok, err := fn(rv.Interface().(T))
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
	}

	switch rv.Kind() {
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i += 1 {
			if !rt.Field(i).IsExported() {
				continue
			}
			if err := travers(rv.Field(i), rt.Field(i).Name, fn); err != nil {
				return err
			}
		}
		return nil
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			if err := travers(rv.MapIndex(key), key.String(), fn); err != nil {
				return err
			}
		}
	case reflect.Slice:
		for i := 0; i < rv.Len(); i += 1 {
			if err := travers(rv.Index(i), strconv.Itoa(i), fn); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		return travers(rv.Elem(), rt.Name(), fn)
	case reflect.Interface:
		return travers(rv.Elem(), rt.Name(), fn)
	}

	return nil

}
