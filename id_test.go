package domain_test

import (
	"testing"

	"github.com/sokool/domain"
)

func TestNewID(t *testing.T) {

	if _, err := domain.NewID[*domain.ID]("asd"); err == nil {
		t.Fatal(err)
	}

	if _, err := domain.NewID[*domain.ID]("86b7cdbb-1c03-409e-a6e9-65501ff97036"); err != nil {
		t.Fatal(err)
	}
}
