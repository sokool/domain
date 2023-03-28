package domain

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var spaces = regexp.MustCompile(`\s+`)

type Name struct {
	First, Last string
	Middle      []string
}

func NewName(text string, maxsize ...int) (Name, error) {
	var n Name
	var l = 3
	var s = spaces.ReplaceAllString(strings.TrimSpace(text), " ")
	var p = strings.Split(s, " ")
	if len(maxsize) >= 1 {
		l = maxsize[0]
	}
	if len(s) < l {
		return n, Errorf("is too short")
	}
	switch m := len(p); {
	case m == 1:
		n.First = p[0]
	case m == 2:
		n.First, n.Last = p[0], p[1]
	default:
		n.First, n.Middle, n.Last = p[0], p[1:m-1], p[m-1]
	}

	n.First, n.Last = strings.Title(n.First), strings.Title(n.Last)
	return n, nil
}

func (f Name) String() string {
	s := fmt.Sprintf("%s %s %s", f.First, strings.Join(f.Middle, " "), f.Last)
	return spaces.ReplaceAllString(strings.TrimSpace(s), " ")
}

func (f Name) IsZero() bool {
	return f.First == "" && f.Last == "" && len(f.Middle) == 0
}

func (f Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.String())
}

func (f *Name) UnmarshalJSON(b []byte) (err error) {
	var s string
	var v Name
	if err = json.Unmarshal(b, &s); err != nil {
		return
	}

	if v, err = NewName(s); err != nil {
		return
	}

	*f = v
	return nil
}

func (f *Name) UnmarshalText(b []byte) error {
	v, err := NewName(string(b))
	if err != nil {
		return err
	}

	*f = v
	return nil
}

func (f *Name) UnmarshalYAML(n *yaml.Node) (err error) {
	var s string
	var v Name

	if err = n.Decode(&s); err != nil {
		return
	}

	if v, err = NewName(s); err != nil {
		return
	}

	*f = v
	return nil
}
