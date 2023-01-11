package domain

import (
	"encoding"
	"github.com/google/uuid"
)

type ID struct{ uuid.UUID }

func (id ID) IsZero() bool {
	return id.UUID == uuid.Nil
}

func NewID[T encoding.TextUnmarshaler](id T, s ...string) (err error) {
	var u uuid.UUID
	var b []byte

	if len(s) == 0 {
		if u, err = uuid.NewRandom(); err != nil {
			return
		}
		if b, err = u.MarshalText(); err != nil {
			return
		}
	} else {
		b = []byte(s[0])
		if u, err = uuid.ParseBytes(b); err != nil {
			return
		}
	}
	return id.UnmarshalText(b)
}
