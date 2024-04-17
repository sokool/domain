package domain

import "testing"

func TestLocation_UnmarshalJSON(t *testing.T) {

}

func TestNewLocation(t *testing.T) {
	cases := []struct {
		name    string
		id      ID[Location]
		point   Point
		address Address
		wantErr bool
	}{
		{
			name:    "Valid Location",
			id:      NewID[Location](),
			point:   MustPoint(37.7749, -122.4194),
			address: Address{Country: "USA"},
			wantErr: false,
		},
		{
			name:    "Invalid ID",
			point:   MustPoint(37.7749, -122.4194),
			address: Address{Country: "PL"},
			wantErr: true,
		},
		{
			name:    "Invalid Point",
			id:      NewID[Location](),
			address: Address{Country: "DE"},
			wantErr: true,
		},
		{
			name:    "Invalid Address",
			id:      NewID[Location](),
			point:   MustPoint(37.7749, -122.4194),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := NewLocation(c.id, c.point, c.address)
			if (err != nil) != c.wantErr {
				t.Errorf("NewLocation() error = %v, wantErr = %v", err, c.wantErr)
				return
			}
		})
	}
}
