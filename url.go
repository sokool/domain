package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type URL struct {
	Schema, Username, Password, Host, Port string
	Path                                   []string
	Query                                  url.Values
	u                                      *url.URL
}

func NewURL(s string) (URL, error) {
	u, err := url.ParseRequestURI(s)
	if err != nil {
		return URL{}, err
	}
	if u.Scheme == "" {
		return URL{}, fmt.Errorf("url: no schema like http or https")
	}
	if u.Host == "" {
		return URL{}, fmt.Errorf("url: host is missing in %s", s)
	}
	var pth []string
	if p := strings.Split(u.Path, "/"); len(p) > 1 {
		pth = p[1:]
	}
	pwd, _ := u.User.Password()
	return URL{
		Schema:   u.Scheme,
		Username: u.User.Username(),
		Password: pwd,
		Host:     u.Hostname(),
		Port:     u.Port(),
		Path:     pth,
		Query:    u.Query(),
		u:        u,
	}, nil
}

func (u *URL) IsZero() bool {
	return u.Schema == "" && u.Host == ""
}

func (u *URL) UnmarshalJSON(b []byte) error {
	var s string
	var err error
	if len(b) == 0 || bytes.Equal(b, []byte(`null`)) {
		return nil
	}
	if err = json.Unmarshal(b, &s); err != nil {
		return err
	}
	*u, err = NewURL(s)
	return err
}

func (u *URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func (u *URL) String() string {
	return u.u.String()
}
