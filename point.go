package domain

import (
	"encoding/json"
	"fmt"
)

type Point struct{ lat, lon *float64 }

func NewPoint(lat, lon float64) (Point, error) {
	if lon < -180 || lon > 180 {
		return Point{}, ErrPoint.New("invalid `%f` longitude", lon)
	}
	if lat < -90 || lat > 90 {
		return Point{}, ErrPoint.New("invalid `%f` latitude", lat)
	}
	return Point{&lat, &lon}, nil
}

func MustPoint(lat, lng float64) Point {
	p, err := NewPoint(lat, lng)
	if err != nil {
		panic(err)
	}
	return p
}

func (p Point) IsEmpty() bool {
	return p.lon == nil || p.lat == nil
}

func (p Point) String() string {
	if p.IsEmpty() {
		return ""
	}
	return fmt.Sprintf("%f,%f", *p.lat, *p.lon)
}

func (p *Point) UnmarshalJSON(bytes []byte) error {
	var b struct{ Lat, Lon *float64 }
	var err error
	if err = json.Unmarshal(bytes, &b); err != nil {
		return ErrPoint.Wrap(err)
	}
	if b.Lat == nil || b.Lon == nil {
		return ErrPoint.New("lat and lng are required")
	}
	if *p, err = NewPoint(*b.Lat, *b.Lon); err != nil {
		return err
	}
	return nil
}

func (p *Point) Scan(data any) error {
	if data == nil {
		return nil
	}
	var lat, lon float64
	switch v := data.(type) {
	case string:
		if _, err := fmt.Sscanf(v, "(%f,%f)", &lat, &lon); err != nil {
			return ErrPoint.Wrap(err)
		}
	default:
		return ErrPoint.New("cannot convert %v to a point", data)
	}
	t, err := NewPoint(lat, lon)
	if err != nil {
		return err
	}
	*p = t
	return nil
}

func (p Point) MarshalJSON() ([]byte, error) {
	if p.IsEmpty() {
		return []byte(`null`), nil
	}
	return []byte(fmt.Sprintf(`{"lat":%v, "lon":%v}`, *p.lat, *p.lon)), nil
}

var ErrPoint = Errorf("point")
