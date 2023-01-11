package domain

import "fmt"

type Error struct{ error }

func Err(space, message string, args ...any) Error {
	message = fmt.Sprintf("%s: %s", space, message)
	return Error{fmt.Errorf(message, args...)}
}
