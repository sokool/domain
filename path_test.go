package domain_test

import (
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
