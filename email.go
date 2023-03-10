package domain

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	Name, Alias, Domain string
}

func NewEmail(text string) (e Email, err error) {
	if ok, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, text); !ok {
		return e, fmt.Errorf("email: %s has invalid format", text)
	}

	p := strings.Split(text, "@")
	e.Name, e.Domain = p[0], p[1]
	if s := strings.Split(p[0], "+"); len(s) != 1 {
		e.Name, e.Alias = s[0], strings.Join(s[1:], "+")
	}

	return e, nil
}

func (e Email) IsZero() bool {
	return e.Name == ""
}

func (e Email) String() string {
	if e.IsZero() {
		return ""
	}
	if e.Alias != "" {
		return fmt.Sprintf("%s+%s@%s", e.Name, e.Alias, e.Domain)
	}
	return fmt.Sprintf("%s@%s", e.Name, e.Domain)
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())

}

func (e *Email) UnmarshalJSON(b []byte) (err error) {
	var s string
	var v Email
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	if v, err = NewEmail(s); err != nil {
		return
	}

	*e = v
	return nil
}

func (e *Email) UnmarshalText(b []byte) error {
	v, err := NewEmail(string(b))
	*e = v
	return err
}
