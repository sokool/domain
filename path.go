package domain

import (
	"regexp"
	"strings"
)

type Path struct{ string }

func NewPath(s string) (Path, error) {
	if s == "/" {
		return Path{s}, nil
	}
	if ok, _ := regexp.MatchString(`^(/[\w\s~+!'.-]+)+$`, s); ok {
		return Path{s}, nil
	}
	return Path{}, Errorf("`%s` must start from / character followed by alphanumerics and/or _ ' ! . - ", s)
}

func (d Path) Append(r Path) Path {
	if d.string == "/" {
		return r
	}
	d.string += r.string
	return d
}

// Cut first n character from underlying dir string
func (d Path) Cut(n int) string {
	if d.IsZero() {
		return ""
	}
	return d.string[n:]
}

// Tail gives last part of dir, it might be directory or file with extension
func (d Path) Tail() string {
	s := strings.Split(d.string, "/")
	return s[len(s)-1]
}

func (d Path) Element(n int) string {
	if d.IsZero() {
		return ""
	}
	s := strings.Split(d.string, "/")
	if len(s)-1 <= n {
		return ""
	}
	return s[n+1]
}

func (d Path) IsZero() bool {
	return d.string == ""
}

func (d Path) String() string {
	return d.string
}
