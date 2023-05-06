package domain

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

type Option uint

const (
	URL_Schema Option = 1 << iota
	URL_Username
	URL_Password
	URL_Host
	URL_Port
	URL_Path
	URL_Query
)

func (o Option) Has(p Option) bool { return o&p != 0 }

type URL struct {
	Schema, Username, Password, Host, Port string
	Path                                   Path
	Query                                  url.Values
	u                                      *url.URL
}

func NewURL(s string, validation ...Option) (URL, error) {
	var u URL
	var v Option
	var ok bool
	for i := range validation {
		v |= validation[i]
	}
	x, err := url.Parse(s)
	if err != nil {
		return u, Errorf("parse failed %w", err)
	}
	if v == 0 {
		v = URL_Schema | URL_Host | URL_Port
	}
	if u.Schema = x.Scheme; v.Has(URL_Schema) && u.Schema == "" {
		return u, Errorf("no schema like http or https")
	}
	if u.Password, ok = x.User.Password(); v.Has(URL_Password) && !ok {
		return u, Errorf("password required")
	}
	if u.Username = x.User.Username(); v.Has(URL_Username) && u.Password == "" {
		return u, Errorf("username required")
	}
	if u.Host = x.Host; v.Has(URL_Host) && u.Host == "" {
		return u, Errorf("host required")
	}
	if u.Port = x.Port(); v.Has(URL_Port) && u.Port == "" {
		return u, Errorf("port required")
	}
	if u.Query = x.Query(); v.Has(URL_Query) && len(u.Query) == 0 {
		return u, Errorf("query required")
	}
	if u.Path, err = NewPath(x.Path); err != nil {
		return URL{}, Errorf("invalid path %w", err)
	}
	return u, nil
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
	*u, err = NewURL(s, 0)
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
