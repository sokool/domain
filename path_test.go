package domain_test

import (
	"fmt"
	"testing"

	"github.com/sokool/domain"
)

func TestNewPath(t *testing.T) {
	type scenario struct {
		description string
		path        string
		err         bool
	}

	cases := []scenario{
		{"root path is required", "/", false},
		{"strings are fine", "/some/cool/files/path/readme", false},
		{"uppercase are ok", "/HI/THERE", false},
		{"numbers are ok", "/users/35682/file", false},
		{"whitespaces are ok", "/Nice Users/Tom Hilldinor", false},
		{"hyphens are ok", "/documents/invoices-2022-04/fv-1", false},
		{"dashes are ok", "/hello_word/_example_/__fold er__", false},
		{"tilda are ok", "/~Hello_word~2~example", false},
		{"dots are ok", "/some.user.path.with/filename.txt", false},
		{"plus are ok", "/in+flames/guitar+samples/ yoo mama", false},
		{"exclamation mark are ok", "/!!important!!!/D.O.C.U.M.E.N.T.S", false},
		{"no slash at beginning is not ok", "some/path/folder", true},
		{"slash at end is not ok", "/some/path/folder/", true},
		{"empty string is not ok", "", true},
		{"dollar sign is not ok", "/$HI$/THERE", true},
		{"multiple slashes are not ok", "/path//name", true},
		{"extra characters at end are not ok", "/dir/game.exe?query=true", true},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			if _, err := domain.NewPath(c.path); (err != nil && !c.err) || (err == nil && c.err) {
				t.Fatalf("expected error:%v got:%v", c.err, err)
			}
		})
	}
}

func TestPath_Element(t *testing.T) {
	type scenario struct {
		description string
		path        string
		from, to    int
		elements    string
	}

	cases := []scenario{
		{"no path", "", 0, -1, ""},
		{"from 0 of / gives /", "/", 0, -1, "/"},
		{"from 0 of /a/b gives /a/b", "/a/b", 0, -1, "/a/b"},
		{"from 0 of /a/b gives /b", "/a/b", 1, -1, "/b"},
		{"from 2 of /a/b gives none", "/a/b", 2, -1, ""},
		{"from 1 to 3 of /a/b/c/d gives /b/c", "/a/b/c/d", 1, 3, "/b/c"},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			p, _ := domain.NewPath(c.path)
			if s := p.Trim(c.from, c.to).String(); s != c.elements {
				t.Fatalf("expected `%s`, got:`%s`", c.elements, s)
			}
			fmt.Println(p.String())
			fmt.Println(p.Replace("/", ""))
		})
	}

}
