package domain_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/sokool/domain"
)

func TestUnit(t *testing.T) {
	u := 34*domain.Tera + 335*domain.Giga + 25*domain.Mega + 38*domain.Kilo + 88
	if _, s := domain.Capacity(u).Values(); s != "TB" {
		t.Fatalf("expected TB got %s", s)
	}
	if _, s := domain.Frequency(u).Values(); s != "THz" {
		t.Fatalf("expected THz got %s", s)
	}
	if _, s := domain.Throughput(u).Values(); s != "Tbps" {
		t.Fatalf("expected THz got %s", s)
	}
	th := domain.Throughput(u)
	if th.String() != "34.335025038088Tbps" {
		fmt.Println(th)
		t.Fatal()
	}
	if th.Kilo() != 3.4335025038088e+10 {
		t.Fatal()
	}
	if th.Mega() != 3.4335025038088e+7 {
		t.Fatal()
	}
	if th.Giga() != 34335.025038088 {
		t.Fatal()
	}
	if th.Tera() != 34.335025038088 {
		t.Fatal()
	}
	if _, u := th.Values(); u != "Tbps" {
		t.Fatal()
	}
	if b, err := th.MarshalText(); string(b) != `34.335025038088Tbps` || err != nil {
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
			err := json.Unmarshal([]byte(c.throughput), &dt)
			if c.err && err == nil {
				t.Fatalf("error expected")
			}
			if !c.err && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
			if (c.throughput == "null" || c.throughput == "") && dt.String() == `0bps` {
				return
			}
			if err == nil && !strings.Contains(c.throughput, dt.String()) {
				t.Fatalf(`expected %s, got "%s"`, c.throughput, dt)
			}
		})
	}
}

func TestPower_UnmarshalJSON(t *testing.T) {
	type scenario struct {
		description string
		power       string
		err         bool
	}

	cases := []scenario{
		{"string without unit->err", `"550.6"`, true},
		{"number without unit->err", `100`, true},
		{"just unit->err", `kW`, true},
		{"unknown unit->err", `40PkW`, true},
		{"100W->ok", `"100W"`, false},
		{"30kW->ok", `"30kW"`, false},
		{"456.314kW->ok", `"456.314kW"`, false},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			var dt domain.Power
			err := json.Unmarshal([]byte(c.power), &dt)
			if c.err && err == nil {
				t.Fatalf("error expected")
			}
			if !c.err && err != nil {
				t.Fatalf("no error expected, got %v", err)
			}
			if (c.power == "null" || c.power == "") && dt.String() == `0bps` {
				return
			}
			if err == nil && !strings.Contains(c.power, dt.String()) {
				t.Fatalf(`expected %s, got "%s"`, c.power, dt)
			}
		})
	}
}

func TestUnit_GoString(t *testing.T) {
	type scenario struct {
		whenUnit   any
		thenString string
	}

	cases := map[string]scenario{
		"137Mbps":                   {137 * domain.Mbps, "137Mbps"},
		"56kHz div 5 = 11.2kHz":     {56 * domain.KHz / 5, "11.2kHz"},
		"1.4GB div 8 = 43.75MB":     {1.4 * domain.GB / 32, "43.75MB"},
		"158kW div 9 = ~17.6kW":     {158 * domain.KW / 9, "~17.56kW"},
		"43Gbps div 3 = ~14.33Gbps": {43 * domain.Gbps / 3, "~14.33Gbps"},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			if s := c.whenUnit.(fmt.GoStringer).GoString(); s != c.thenString {
				t.Errorf("got %s, want %s", s, c.thenString)
			}
		})
	}
}
