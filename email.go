package domain

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type Email struct {
	Name, Alias, Domain string
}

func NewEmail(text string, re ...*regexp.Regexp) (e Email, err error) {
	if len(re) == 0 {
		re = append(re, email)
	}
	if !re[0].MatchString(text) {
		return e, fmt.Errorf("email: %s has invalid format", text)
	}

	p := strings.Split(text, "@")
	e.Name, e.Domain = p[0], p[1]
	if s := strings.Split(p[0], "+"); len(s) != 1 {
		e.Name, e.Alias = s[0], strings.Join(s[1:], "+")
	}

	return e, nil
}

func (e Email) IsEmpty() bool {
	return e.Name == ""
}

func (e Email) String() string {
	if e.IsEmpty() {
		return ""
	}
	if e.Alias != "" {
		return fmt.Sprintf("%s+%s@%s", e.Name, e.Alias, e.Domain)
	}
	return fmt.Sprintf("%s@%s", e.Name, e.Domain)
}

func (e Email) Hash() string {
	return Hash(e)
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *Email) UnmarshalJSON(b []byte) (err error) {
	var s string
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	*e, err = NewEmail(s)
	return nil
}

func (e *Email) UnmarshalText(b []byte) error {
	v, err := NewEmail(string(b))
	*e = v
	return err
}

func (e *Email) UnmarshalYAML(n *yaml.Node) (err error) {
	var s string
	var v Email

	if err = n.Decode(&s); err != nil {
		return
	}

	if v, err = NewEmail(s); err != nil {
		return
	}

	*e = v
	return nil
}

var email = regexp.MustCompile(`^[\w._%+\-]+@[\w.\-]+\.\w{2,32}$`)
