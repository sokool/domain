package domain

import (
	"regexp"
	"strings"
)

type Path struct{ strings []string }

func NewPath(s string) (Path, error) {
	if len(s) == 0 || s[0] != '/' {
		s = "/" + s
	}
	ss := strings.Split(s, "/")
	if s == "/" {
		return Path{ss}, nil
	}
	if ok, _ := regexp.MatchString(`^(/[\w\s~+!'.-]+)+$`, s); ok {
		return Path{ss}, nil
	}
	return Path{}, Errorf("`%s` must start from / character followed by alphanumerics and/or _ ' ! . - ", s)
}

func (d Path) Append(r Path) Path {
	if r.IsZero() {
		return d
	}
	if d.Trim(0).String() == "/" {
		return r
	}
	return Path{append(d.strings, r.strings[1:]...)}
}

// Replace first n character from underlying dir string
func (d Path) Replace(old, new string, occurrences ...int) string {
	if len(occurrences) == 0 {
		occurrences = append(occurrences, 1)
	}
	return strings.Replace(d.String(), old, new, occurrences[0])
}

// Tail gives last part of dir, it might be directory or file with extension
func (d Path) Tail() string {
	return d.strings[len(d.strings)-1]
}

// Trim ...
func (d Path) Trim(from int, to ...int) Path {
	s := d.Size()
	if s == 0 {
		return Path{}
	}
	if s <= from {
		return Path{}
	}
	ss := []string{""}
	if l := len(to); l == 0 || to[0] == -1 {
		ss = append(ss, d.strings[from+1:]...)
	} else {
		ss = append(ss, d.strings[from+1:to[0]+1]...)
	}
	return Path{ss}
}

func (d Path) Size() int {
	return len(d.strings) - 1
}

func (d Path) IsZero() bool {
	return d.Size() == 0
}

func (d Path) String() string {
	return strings.Join(d.strings, "/")
}
