package domain_test

import (
	"fmt"
	"testing"

	"github.com/sokool/domain"
)

func TestErr(t *testing.T) {
	type scenario struct {
		description                 string
		err                         *domain.Error // given
		name, code, message, string string        // expects
	}

	cases := []scenario{
		{
			description: "empty",
			err:         domain.Errorf(""),
			string:      "github.com/sokool/domain_test.TestErr#error_test.go@L20",
			name:        "github.com/sokool/domain_test.TestErr",
			code:        "error_test.go@L20",
		},
		{
			description: "just message",
			err:         domain.Errorf("hi there"),
			string:      "github.com/sokool/domain_test.TestErr#error_test.go@L27: hi there",
			name:        "github.com/sokool/domain_test.TestErr",
			code:        "error_test.go@L27",
			message:     "hi there",
		},
		{
			description: "message with arguments",
			err:         domain.Errorf("hi there %s", "man"),
			string:      "github.com/sokool/domain_test.TestErr#error_test.go@L35: hi there man",
			name:        "github.com/sokool/domain_test.TestErr",
			code:        "error_test.go@L35",
			message:     "hi there man",
		},
		{
			description: "just name",
			err:         domain.Errorf("test:"),
			string:      "test:",
			name:        "test",
		},
		{
			description: "just code",
			err:         domain.Errorf("#h6b7:"),
			string:      "#h6b7",
			code:        "h6b7",
		},
		{
			description: "with name,message",
			err:         domain.Errorf("test:hi there"),
			string:      "test: hi there",
			name:        "test",
			message:     "hi there",
		},
		{
			description: "with name,code,message",
			err:         domain.Errorf("email#h1:     invalid hostname      "),
			string:      "email#h1: invalid hostname",
			name:        "email",
			code:        "h1",
			message:     "invalid hostname",
		},
		{
			description: "with name,code,message from arguments",
			err:         domain.Errorf("%s#%s: invalid %s", "email", "h45", "username"),
			string:      "email#h45: invalid username",
			name:        "email",
			code:        "h45",
			message:     "invalid username",
		},
		{
			description: "with code,message",
			err:         domain.Errorf("#e87:failed"),
			string:      "#e87: failed",
			code:        "e87",
			message:     "failed",
		},
		{
			description: "with code,name signs in wrong place",
			err:         domain.Errorf("failed due abc: #triggered"),
			name:        "github.com/sokool/domain_test.TestErr",
			string:      "github.com/sokool/domain_test.TestErr#error_test.go@L85: failed due abc: #triggered",
			code:        "error_test.go@L85",
			message:     "failed due abc: #triggered",
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			if s := c.err.Name(); s != c.name {
				t.Fatalf("expected name `%s`, got `%s`", c.name, s)
			}
			if s := c.err.Code(); s != c.code {
				t.Fatalf("expected code `%s`, got `%s`", c.code, s)
			}
			if s := c.err.Message(); s != c.message {
				t.Fatalf("expected message `%s`, got `%s`", c.message, s)
			}
			if s := c.err.Error(); s != c.string {
				t.Fatalf("expected string `%s`, got `%s`", c.string, s)
			}
		})
	}
}

func TestGetErr(t *testing.T) {
	_, a := domain.NewText("dupa", 8, 16)
	b := fmt.Errorf("second %w", a)
	c := fmt.Errorf("third %w", b)
	d := domain.Err(c)

	if a != d {
		t.Fatal()
	}

	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(d)

}
