package domain

import (
	"testing"
)

func TestID_UnmarshalText(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		isError bool
	}{
		{
			name:    "Valid UUID",
			input:   "f47ac10b-58cc-0372-8567-0e02b2c3d479",
			isError: false,
		},
		{
			name:    "Invalid UUID",
			input:   "f47ac10b-58cc-0372-8567-0e02b2c3d47",
			isError: true,
		},
		{
			name:    "Invalid Text",
			input:   "some-id",
			isError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var id ID[string]

			err := id.UnmarshalText([]byte(tt.input))

			if tt.isError && err == nil {
				t.Fatalf("error expected")
			}
			if !tt.isError && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
			if err != nil {
				return
			}
			if id.String() != tt.input {
				t.Errorf("expected id %s, but got %s", tt.input, id)
			}
			if id.IsEmpty() {
				t.Errorf("expected id to not be empty")
			}
		})
	}
}

func TestID_MarshalText(t *testing.T) {
	tests := []struct {
		name     string
		input    ID[string]
		expected string
	}{
		{
			name:     "Valid UUID",
			input:    MustID[string]("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expected: "f47ac10b-58cc-0372-8567-0e02b2c3d479",
		},
		{
			name:     "Empty ID",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b, err = tt.input.MarshalText()

			if err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
			if string(b) != tt.expected {
				t.Errorf("expected id %s, but got %s", tt.expected, string(b))
			}
		})
	}
}
