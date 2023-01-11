package view

import (
	"bytes"
	json "encoding/json"
	"fmt"
	yaml "github.com/goccy/go-yaml"
	"reflect"
)

type Body any

type Renderer interface {
	Render(types ...string) any
}

func String(r Body) string {
	b, err := JSON(r)
	if err != nil {
		return fmt.Sprintf("view: %T to string failed due %s", r, err)
	}
	w := bytes.Buffer{}
	if err = json.Indent(&w, b, "", "\t"); err != nil {
		return fmt.Sprintf("view: %T to string failed due %s", r, err)
	}
	return w.String()
}

func Print(r Body, typ ...string) {
	fmt.Println(String(r))
}

var renderer = reflect.TypeOf((*Renderer)(nil)).Elem()

var rawJSON = reflect.TypeOf((*json.RawMessage)(nil)).Elem()

func JSON(b Body, n ...string) (Bytes, error) {
	switch v := reflect.ValueOf(b); v.Kind() {
	case reflect.Slice:
		if !v.Type().Elem().Implements(renderer) {
			return json.Marshal(b)
		}
		var s = make([]json.RawMessage, v.Len())
		for i := 0; i < v.Len(); i++ {
			m, err := JSON(v.Index(i).Interface(), n...)
			if err != nil {
				return nil, err
			}

			s[i] = m.JSON()
		}

		if len(s) == 0 {
			return Null, nil
		}
		return json.Marshal(s)
	case reflect.Map:
		if v.Len() == 0 {
			return Null, nil
		}
		if !v.Type().Elem().Implements(renderer) && !v.Type().Implements(renderer) {
			return json.Marshal(b)
		}

		f := reflect.MakeMapWithSize(reflect.MapOf(v.Type().Key(), rawJSON), v.Len())
		for m := v.MapRange(); m.Next(); {
			j, err := JSON(m.Value().Interface(), n...)
			if err != nil {
				return nil, err
			}
			f.SetMapIndex(m.Key(), reflect.ValueOf(j.JSON()))
		}
		return json.Marshal(f.Interface())
	}

	if r, ok := b.(Renderer); ok {
		x, err := JSON(r.Render(n...), n...)
		if err != nil {
			return nil, err
		}
		return json.Marshal(x.JSON())
	}

	return json.Marshal(b)
}

func YAML(b Body) (Bytes, error) {
	return yaml.Marshal(b)
}
