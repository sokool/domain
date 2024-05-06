package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Location describes geographical placement
type Location struct {
	ID      ID[Location] `json:"id"`
	Address Address      `json:"address"`
	Point   Point        `json:"point"`
	Meta    Meta         `json:"meta"`
}

func NewLocation(id ID[Location], p Point, a Address, m ...Meta) (Location, error) {
	var r Location
	if r.ID = id; r.ID.IsEmpty() {
		return Location{}, ErrLocation.New("invalid id")
	}
	if r.Point = p; r.Point.IsEmpty() {
		return Location{}, ErrLocation.New("invalid point")
	}
	if r.Address = a; r.Address.IsEmpty() {
		return Location{}, ErrLocation.New("invalid address")
	}
	if len(m) != 0 {
		r.Meta = m[0]
	}
	return r, nil
}

func MustLocation(id ID[Location], p Point, a Address, m ...Meta) Location {
	l, err := NewLocation(id, p, a, m...)
	if err != nil {
		panic(err)
	}
	return l
}

func FakeLocation() Location {
	id := NewID[Location]()
	pt := MustPoint(37.7749, -122.4194)
	ar := Address{Country: "USA", Name: "JFK5"}
	return MustLocation(id, pt, ar)
}

func (p Location) String() string {
	return fmt.Sprintf("%s %s %s", p.ID, p.Address, p.Point)
}

func (p Location) GoString() string {
	b, _ := json.MarshalIndent(p, "", "\t")
	return fmt.Sprintf("%T\n%s\n", p, b)
}

func (p Location) IsEmpty() bool {
	return p.Point.IsEmpty()
}

func (p Location) ExternalID(provider string) string {
	id, _ := p.Meta[fmt.Sprintf("%s_id", strings.ToLower(provider))].(string)
	return id
}

func (p Location) SetExternalID(provider, id string) Location {
	if p.Meta == nil {
		p.Meta = make(Meta)
	}
	p.Meta[fmt.Sprintf("%s_id", provider)] = id
	return p
}

func (p *Location) UnmarshalJSON(b []byte) error {
	type no Location
	var n no
	var err error
	if err = json.Unmarshal(b, &n); err != nil {
		return ErrLocation.Wrap(err)
	}
	if *p, err = NewLocation(n.ID, n.Point, n.Address, n.Meta); err != nil {
		return err
	}
	return nil
}

type Address struct {
	Name    string
	Address string
	Zip     string
	City    string
	State   string
	Country string // ISO3
}

func (a Address) String() string {
	s := fmt.Sprintf(" %s, %s, %s, %s %s, %s", a.Name, a.Address, a.City, a.State, a.Zip, a.Country)
	s = strings.TrimSpace(strings.ReplaceAll(s, " ,", " "))
	if n := len(s); n > 0 && s[n-1] == ',' {
		s = s[:n-1]
	}
	return strings.Join(strings.Fields(s), " ")
}

func (a Address) IsEmpty() bool {
	return a.Country == ""
}

var ErrLocation = Errorf("location")
