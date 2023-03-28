package domain

import (
	"encoding"

	"github.com/google/uuid"
)

type ID struct{ uuid.UUID }

func (id *ID) IsEmpty() bool {
	return id.UUID == uuid.Nil
}

func NewID[T encoding.TextUnmarshaler](t T, s ...string) (err error) {
	var u uuid.UUID
	var b []byte

	if len(s) == 0 {
		if u, err = uuid.NewRandom(); err != nil {
			return Errorf("%w", err)
		}
		if b, err = u.MarshalText(); err != nil {
			return Errorf("%w", err)
		}
	} else {
		b = []byte(s[0])
		if u, err = uuid.ParseBytes(b); err != nil {
			return Errorf("%w", err)
		}
	}

	if err = t.UnmarshalText(b); err != nil {
		return Errorf("%w", err)
	}

	return nil
}
