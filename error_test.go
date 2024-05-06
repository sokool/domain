package domain_test

import (
	"errors"
	"testing"

	"github.com/sokool/domain"
)

func TestErr(t *testing.T) {
	var (
		Err1 = domain.Errorf("one")
		Err2 = domain.Errorf("%w:two", Err1)
		Err3 = domain.Errorf("three")
		Err4 = domain.Errorf("%w: oh no", Err2)
		Err5 = Err2.New("hi %w", Err3)
	)

	if errors.Is(Err4, Err3) {
		t.Fatal()
	}
	if !errors.Is(Err4, Err2) {
		t.Fatal()
	}
	if !errors.Is(Err4, Err1) {
		t.Fatal()
	}
	if !errors.Is(Err5, Err1) {
		t.Fatal()
	}
	if !errors.Is(Err5, Err2) {
		t.Fatal()
	}
	if !errors.Is(Err5, Err3) {
		t.Fatal()
	}
}
