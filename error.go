package domain

import (
	"errors"
	"fmt"
)

type Error struct {
	message string
	wrapped error
}

func Errorf(msg string, args ...any) *Error {
	err := fmt.Errorf(msg, args...)
	return &Error{err.Error(), errors.Unwrap(err)}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) New(msg string, args ...any) error {
	err := Errorf(msg, args...)
	return &Error{e.message + ":" + err.message, fmt.Errorf("%w %w", e, err)}
}

func (e *Error) Wrap(err error) error {
	return e.New("%w", err)
}

func (e *Error) Unwrap() error { return e.wrapped }
