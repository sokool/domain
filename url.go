package domain

import (
	"bytes"
	"encoding/json"
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
		return URL{}, Errorf("%w", err)
	}
	if u.Scheme == "" {
		return URL{}, Errorf("no schema like http or https")
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

func (u *URL) Format(s string) string {
	s = strings.Replace(s, "scheme", u.Schema, -1)
	s = strings.Replace(s, "host", u.Host, -1)
	s = strings.Replace(s, "port", u.Port, -1)
	s = strings.Replace(s, "user", u.Username, -1)
	s = strings.Replace(s, "password", u.Password, -1)
	s = strings.Replace(s, "path", strings.Join(u.Path, "/"), -1)
	s = strings.Replace(s, "query", u.Query.Encode(), -1)
	return s
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
