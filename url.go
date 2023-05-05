package domain

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

type URL struct {
	Schema, Username, Password, Host, Port string
	Path                                   Path
	Query                                  url.Values
	u                                      *url.URL
}

func NewURL(s string, validation ...string) (URL, error) {
	v := texts(validation)
	u, err := url.Parse(s)
	if err != nil {
		return URL{}, Errorf("parse failed %w", err)
	}
	if s == "" && v.has("no-empty") == -1 {
		return URL{}, Errorf("no empty string allowed")
	}
	if u.Scheme == "" && v.has("schema") != -1 {
		return URL{}, Errorf("no schema like http or https")
	}

	var d Path
	if p := strings.Split(u.Path, "/"); len(p) > 1 {
		if d, err = NewPath(u.Path); err != nil {
			return URL{}, Errorf("invalid path %w", err)
		}
	}
	pwd, _ := u.User.Password()
	return URL{
		Schema:   u.Scheme,
		Username: u.User.Username(),
		Password: pwd,
		Host:     u.Hostname(),
		Port:     u.Port(),
		Path:     d,
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
	s = strings.Replace(s, "path", u.Path.String(), -1)
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

type texts []string

func (s texts) has(value string) int {
	for i := range s {
		if s[i] == value {
			return i
		}
	}
	return -1
}
