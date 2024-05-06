package domain

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type ID[T any] string

// NewID creates a new ID[T] from a random UUID.
func NewID[T any]() ID[T] {
	return ID[T](uuid.New().String())
}

// ParseID creates a new ID[T] from a string or returns an error if the string is not a valid UUID.
func ParseID[T any](s string) (ID[T], error) {
	s = strings.TrimSpace(s)
	if _, err := uuid.Parse(s); err != nil {
		var v T
		return "", ErrID.New("%T: %s", v, err)
	}
	return ID[T](s), nil
}

// MustID creates a new ID[T] from a string or panics if the string is not a valid UUID.
func MustID[T any](s string) ID[T] {
	var t ID[T]
	var err error
	if t, err = ParseID[T](s); err != nil {
		panic(err)
	}
	return t
}

// Hash creates a new hashed UUID from any given list of strings.
// The strings are concatenated and hashed using SHA1.
func Hash(s ...fmt.Stringer) string {
	var k []byte
	for i := range s {
		k = append(k, s[i].String()...)
	}
	return uuid.NewSHA1(uuid.NameSpaceDNS, k).String()
}

func (id ID[T]) IsEmpty() bool {
	return id == ""
}

func (id ID[T]) String() string {
	return string(id)
}

func (id ID[T]) MarshalText() ([]byte, error) {
	return []byte(id), nil
}

func (id *ID[T]) UnmarshalText(b []byte) error {
	var err error
	if *id, err = ParseID[T](string(b)); err != nil {
		return err
	}
	return nil
}

var ErrID = Errorf("id")
