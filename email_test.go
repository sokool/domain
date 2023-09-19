package domain_test

import (
	"testing"

	"github.com/sokool/domain"
)

func TestNewEmail(t *testing.T) {
	type scenario struct {
		description string
		mail        string
		err         bool
	}

	cases := []scenario{
		{"empty->err", "", true},
		{"@b.com->err", "@b.com", true},
		{"a@->err", "a@", true},
		{"a@b->err", "a@b", true},
		{"a@b.c->err", "a@b.c", true},
		{"a@b.us->ok", "a@b.us", false},
		{"Brad.Girow4+rollee+test@You2.CAN-Win.web3->ok", "Brad.Girow4+rollee+test@You2.CAN-Win.web3", false},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			_, err := domain.NewEmail(c.mail)
			if c.err && err == nil {
				t.Fatalf("error expected")
			}
			if !c.err && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
		})
	}
}
