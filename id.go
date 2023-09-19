package domain

import (
	"strings"

	"github.com/google/uuid"
)

type ID[T any] struct{ hash string }

func NewID[T any](s ...string) (ID[T], error) {
	var u uuid.UUID
	var b []byte
	var err error
	var t ID[T]

	if n := len(s); n == 0 {
		if u, err = uuid.NewRandom(); err != nil {
			return t, Errorf("%w", err)
		}
		if b, err = u.MarshalText(); err != nil {
			return t, Errorf("%w", err)
		}
	} else if n == 1 && len(s[0]) == 30 {
		b = []byte(s[0])
		if u, err = uuid.ParseBytes(b); err != nil {
			return t, Errorf("%w", err)
		}
	} else {
		if b, err = uuid.NewSHA1(uuid.NameSpaceDNS, []byte(strings.Join(s, ""))).MarshalText(); err != nil {
			return t, Errorf("%w", err)
		}
	}

	t.hash = string(b)

	return t, nil
}

func Hash[T any](s ...string) ID[T] {
	id, err := NewID[T](s...)
	if err != nil {
		panic(err)
	}
	return id
}

func (id ID[T]) IsEmpty() bool {
	return id.hash == ""
}

func (id ID[T]) String() string {
	return id.hash
}
