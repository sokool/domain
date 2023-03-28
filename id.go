package domain

import (
	"encoding"
	"reflect"

	"github.com/google/uuid"
)

type ID struct{ uuid.UUID }

func (id *ID) IsEmpty() bool {
	return id.UUID == uuid.Nil
}

func NewID[T encoding.TextUnmarshaler](s ...string) (t T, err error) {
	var u uuid.UUID
	var b []byte

	if len(s) == 0 {
		if u, err = uuid.NewRandom(); err != nil {
			return t, Errorf("%w", err)
		}
		if b, err = u.MarshalText(); err != nil {
			return t, Errorf("%w", err)
		}
	} else {
		b = []byte(s[0])
		if u, err = uuid.ParseBytes(b); err != nil {
			return t, Errorf("%w", err)
		}
	}

	t = reflect.New(reflect.TypeOf(t).Elem()).Interface().(T)
	if err = t.UnmarshalText(b); err != nil {
		return t, Errorf("%w", err)
	}

	return t, nil
}
