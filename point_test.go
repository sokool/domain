package domain

import (
	"encoding/json"
	"testing"
)

func TestNewPoint(t *testing.T) {
	tests := []struct {
		name     string
		lat      float64
		lng      float64
		expected Point
		isError  bool
	}{
		{
			name:     "Valid Point",
			lat:      40.7128,
			lng:      -74.0060,
			expected: MustPoint(40.7128, -74.0060),
			isError:  false,
		},
		{
			name:     "Invalid Latitude",
			lat:      100.0,
			lng:      -74.0060,
			expected: Point{},
			isError:  true,
		},
		{
			name:     "Invalid Longitude",
			lat:      40.7128,
			lng:      -200.0,
			expected: Point{},
			isError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewPoint(tt.lat, tt.lng)

			if tt.isError && err == nil {
				t.Fatalf("error expected")
			}
			if !tt.isError && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
			if err != nil {
				return
			}
			if p.String() != tt.expected.String() {
				t.Errorf("Expected point %v, but got %v", tt.expected, p)
			}
		})
	}
}

func TestPoint_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Point
		isError  bool
	}{
		{
			name:     "Valid JSON",
			input:    `{"lat": 40.7128, "lon": -74.0060}`,
			expected: MustPoint(40.7128, -74.0060),
			isError:  false,
		},
		{
			name:     "Zero Latitude and Longitude is Valid JSON",
			input:    `{"lat": 0, "lon": 0}`,
			expected: MustPoint(0, 0),
			isError:  false,
		},
		{
			name:     "Invalid JSON",
			input:    `{"lat": "invalid", "lon": -74.0060}`,
			expected: Point{},
			isError:  true,
		},
		{
			name:     "Missing Latitude",
			input:    `{"lon": -74.0060}`,
			expected: Point{},
			isError:  true,
		},
		{
			name:     "Missing Longitude",
			input:    `{"lat": 40.7128}`,
			expected: Point{},
			isError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Point
			var err = json.Unmarshal([]byte(tt.input), &p)
			if tt.isError && err == nil {
				t.Fatalf("error expected")
			}
			if !tt.isError && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
		})
	}
}

func TestPoint_Scan(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected Point
		isError  bool
	}{
		{
			name:     "Valid Point",
			input:    "(40.7128,-74.0060)",
			expected: MustPoint(40.7128, -74.0060),
			isError:  false,
		},
		{
			name:     "Invalid Point",
			input:    "invalid",
			expected: Point{},
			isError:  true,
		},
		{
			name:     "Nil",
			input:    nil,
			expected: Point{},
			isError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var p Point
			var err = p.Scan(tt.input)
			if tt.isError && err == nil {
				t.Fatalf("error expected")
			}
			if !tt.isError && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
		})
	}
}
