package domain_test

import (
	"fmt"
	"testing"

	"github.com/sokool/domain"
)

func TestURL_UnmarshalJSON(t *testing.T) {
	//var u domain.URL
	//assert.Error(t, u.UnmarshalJSON([]byte(`""`)))
	//assert.Error(t, u.UnmarshalJSON([]byte(`"       asdf r"`)))
	//assert.Error(t, u.UnmarshalJSON([]byte(`"rollee"`)))
	//
	//assert.NoError(t, u.UnmarshalJSON(nil))
	//assert.NoError(t, u.UnmarshalJSON([]byte(``)))
	//assert.NoError(t, u.UnmarshalJSON([]byte(`null`)))
	//assert.NoError(t, u.UnmarshalJSON([]byte(`"https://greg:j6G1Df7@rollee.io:8888/documents/january-invoice.pdf?version=3"`)))
	//
	//assert.Equal(t, u.Schema, "https")
	//assert.Equal(t, u.Username, "greg")
	//assert.Equal(t, u.Password, "j6G1Df7")
	//assert.Equal(t, u.Host, "rollee.io")
	//assert.Equal(t, u.Port, "8888")
	//assert.Equal(t, u.Path, []string{"documents", "january-invoice.pdf"})
	//assert.Equal(t, u.Query, url.Values{"version": {"3"}})
}

func TestURL_IsZero(t *testing.T) {
	//var u domain.URL
	//assert.Equal(t, u.IsZero(), true)
}

func TestNewURL(t *testing.T) {
	type scenario struct {
		description string
		url         string
		err         bool
	}

	cases := []scenario{
		{"empty string fails", "", true},
		{"host, no schema ok", "wosp.org.pl", false},
		{"schema,host is ok", "https://wosp.org.pl", false},
		{"schema,host,path is ok", "https://wosp.org.pl/some/path", false},
	}

	u, err := domain.NewURL("https://wosp.org.pl/some/path")
	fmt.Println(u.Path, err)
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			if _, err := domain.NewURL(c.url); (err != nil && !c.err) || (err == nil && c.err) {
				t.Fatalf("expected error:%v got:%v", c.err, err)
			}
		})
	}
}
