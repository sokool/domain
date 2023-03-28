package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Error struct {
	s struct{ Message, Name, Code string }
}

func NewError(message, name, code string, caller ...int) *Error {
	if len(caller) == 0 {
		caller = append(caller, 1)
	}
	var err Error
	message, name, code = strings.TrimSpace(message), strings.TrimSpace(name), strings.TrimSpace(code)
	if name == "" && code == "" {
		if p, f, l, ok := runtime.Caller(caller[0]); ok {
			name, code = runtime.FuncForPC(p).Name(), fmt.Sprintf("%s@L%d", f[strings.LastIndex(f, "/")+1:], l)
		}
	}
	err.s.Message, err.s.Name, err.s.Code = message, name, code
	return &err
}

func Errorf(format string, args ...any) *Error {
	var m, n, c = fmt.Errorf(format, args...).Error(), "", ""
	if i := strings.Index(m, ":"); i > 0 && !strings.Contains(m[:i], " ") {
		n, m = m[:i], m[i+1:]
		if k := strings.Index(n, "#"); k >= 0 {
			c, n = n[k+1:], n[:k]
		}
	}
	return NewError(m, n, c, 2)
}

func Err(from error) *Error {
	var e *Error
	if errors.As(from, &e) {
		return e
	}
	return nil
}
func (e *Error) Name() string {
	return e.s.Name
}

func (e *Error) Code() string {
	return e.s.Code
}

func (e *Error) Message() string {
	return e.s.Message
}

func (e *Error) Error() string {
	s := fmt.Sprintf("%s#%s: %s", e.s.Name, e.s.Code, e.s.Message)
	if s[:3] == "#: " {
		s = s[3:]
	}
	if n := len(s); n > 2 && s[n-2:] == ": " {
		s = s[:n-2]
	}
	if n := len(s); n > 1 && s[n-1:] == "#" {
		s = s[:n-1] + ":"
	}
	return strings.Replace(s, "#:", ":", -1)
}

func (e *Error) Unwrap() error {
	fmt.Println("unwrapping")
	return nil
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.s)
}
