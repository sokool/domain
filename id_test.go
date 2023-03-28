package domain_test

import (
	"testing"

	"github.com/sokool/domain"
)

func TestNewID(t *testing.T) {
	var id domain.ID
	if err := domain.NewID(&id, "asd"); err == nil {
		t.Fatal(err)
	}

	if err := domain.NewID(&id, "86b7cdbb-1c03-409e-a6e9-65501ff97036"); err != nil {
		t.Fatal(err)
	}
}
