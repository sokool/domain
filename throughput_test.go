package domain_test

import (
	"strings"
	"testing"

	"github.com/sokool/domain"
)

func TestThroughput(t *testing.T) {
	th := 34*domain.Tbps + 335*domain.Gbps + 25*domain.Mbps + 38*domain.Kbps + 88*domain.Bps
	if th.String() != "34.335025038088Tbps" {
		t.Fatal()
	}
	if th.Bits() != 3.4335025038088e+13 {
		t.Fatal()
	}
	if th.Kilobits() != 3.4335025038088e+10 {
		t.Fatal()
	}
	if th.Megabits() != 3.4335025038088e+7 {
		t.Fatal()
	}
	if th.Gigabits() != 34335.025038088 {
		t.Fatal()
	}
	if th.Terabits() != 34.335025038088 {
		t.Fatal()
	}
	if b, err := th.MarshalJSON(); string(b) != `"34.335025038088Tbps"` || err != nil {
		t.Fatal()
	}
}

func TestThroughput_UnmarshalJSON(t *testing.T) {
	type scenario struct {
		description string
		throughput  string
		err         bool
	}

	cases := []scenario{
		{"empty->err", ``, true},
		{"string without unit->err", `"550.6"`, true},
		{"number without unit->err", `100`, true},
		{"just unit->err", `Mbps`, true},
		{"unknown unit->err", `40Pbps`, true},
		{"null->ok", `null`, false},
		{"30bps->ok", `"30bps"`, false},
		{"456.314kbps->ok", `"456.314kbps"`, false},
		{"757Mbps->ok", `"757Mbps"`, false},
		{"10Gbps->ok", `"10Gbps"`, false},
		{"3.45Tbps->ok", `"3.45Tbps"`, false},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			var dt domain.Throughput
			var err = dt.UnmarshalJSON([]byte(c.throughput))
			if c.err && err == nil {
				t.Fatalf("error expected")
			}
			if !c.err && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
			if c.throughput == "null" && dt.String() == `0bps` {
				return
			}
			if err == nil && !strings.Contains(c.throughput, dt.String()) {
				t.Fatalf(`expected %s, got "%s"`, c.throughput, dt)
			}
		})
	}
}
