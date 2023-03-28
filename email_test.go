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
		{"empty not ok", "", true},
		{"@b.com not ok", "@b.com", true},
		{"a@ not ok", "a@", true},
		{"a@b not ok", "a@b", true},
		{"a@b.c not ok", "a@b.c", true},
		{"a@b.us->ok", "a@b.us", false},
		{"Brad.Girow4+rollee+test@You2.CAN-Win.web3 is ok", "Brad.Girow4+rollee+test@You2.CAN-Win.web3", false},
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
