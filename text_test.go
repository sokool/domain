package domain_test

import (
	"github.com/sokool/domain"
	"testing"
)

func TestBodyJson(t *testing.T) {
	n := domain.Text("hi there")

	if n.IsZero() {
		t.Fatalf("expected not zero, got %s", n)
	}

	if n.Words() != 2 {
		t.Fatal()
	}

	if n.Word(0) != "hi" || n.Word(1) != "there" {
		t.Fatal("x")
	}

	if s := n.Length(); s != 8 {
		t.Fatalf("wrong size, got %d", s)
	}

	if !n.LengthBetween(0, n.Length()) {
		t.Fatalf("wrong length")
	}

	b, err := n.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	var m domain.Text
	if err = m.UnmarshalJSON(b); err != nil {
		t.Fatal(err)
	}

	if n != m {
		t.Fatalf("expected %s, got  %s", n, m)
	}
}

func TestJSONMultiUnmarshal(t *testing.T) {
	n := domain.Text("Hello\nWorld")
	for i := 0; i < 5; i++ {
		x, err := n.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}

		if err = n.UnmarshalJSON(x); err != nil {
			t.Fatal(err)
		}
	}

	if n != "Hello\nWorld" {
		t.Fatal("uuu")
	}
}
