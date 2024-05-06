package domain

import (
	"encoding/json"
	"testing"
)

func TestMoney(t *testing.T) {
	m := MustMoney("Snickers: -3.89956 USD")
	if f := m.Float(); f != -3.89 {
		t.Fatalf("expected -3.89, got %g", f)
	}
	if m.IsEmpty() {
		t.Fatal()
	}
	if s := m.Amount; *s != -389 {
		t.Fatalf("expected -389, got %d", *s)
	}
	if s := m.Amount.String(); s != "-3.89" {
		t.Fatalf("expected -3.89, got %s", s)
	}
	if s := m.Currency; s != "USD" {
		t.Fatalf("expected USD, got %s", s)
	}
	if s := m.Label; s != "Snickers" {
		t.Fatalf("expected Snickers, got %s", s)
	}
	if m.String() != "Snickers: -3.89 USD" {
		t.Fatal()
	}
	if m, _ = m.Add(MustMoney("4.188888 USD")); m.Float() != 0.29 {
		t.Fatal()
	}
}

func TestParseMoney(t *testing.T) {
	type scenario struct {
		// given
		in string
		// expects
		err    bool
		amount string
		label  string
	}

	cases := map[string]scenario{
		"empty fails":               {in: "", err: true},
		"random text fails":         {in: "something ", err: true},
		"just amount fails":         {in: "2", err: true},
		"amount and currency is ok": {in: "2 GBP", err: false, amount: "2.00"},
		"negative amount is ok":     {in: "-12 PLN", err: false, amount: "-12.00"},
		//"floating point amount is ok": {in: "32.48 NZD", err: false, amount: "32.48"},
		"rounding 3.191 to 3.19": {in: "3.191 USD", err: false, amount: "3.19"},
		"rounding 3.195 to 3.19": {in: "3.195 USD", err: false, amount: "3.19"},
		"rounding 3.199 to 3.19": {in: "3.199 USD", err: false, amount: "3.19"},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			m, err := ParseMoney(c.in)
			if c.err && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !c.err && err != nil {
				t.Fatalf("expected no error, got %s", err)
			}
			if err != nil {
				return
			}
			if c.amount != m.Amount.String() {
				t.Fatalf("expected %s, got %s", c.amount, m.Amount.String())
			}
			if c.label != m.Label {
				t.Fatalf("expected %s, got %s", c.label, m.Label)
			}
		})
	}
}

func TestMoney_UnmarshalJSON(t *testing.T) {
	type scenario struct {
		description string
		input       string
		output      string
	}

	cases := []scenario{
		{
			"empty",
			``,
			`null`,
		},
		{
			"null",
			`null`,
			`null`,
		},
		{
			"just a number is not ok",
			`1`,
			`money:invalid '1' string format`,
		},
		{
			"just a currency is not ok",
			`"PLN"`,
			`money:invalid 'PLN' string format`,
		},
		{
			"currency and value as string is ok",
			`"1 USD"`,
			`{"amount":1.00,"currency":"USD"}`,
		},
		{
			"negative amount with label",
			`"Spaghetti: -5.99EUR"`,
			`{"amount":-5.99,"currency":"EUR","label":"Spaghetti"}`,
		},
		{
			"rounding down",
			`"3.49000000000004 GBP"`,
			`{"amount":3.49,"currency":"GBP"}`,
		},
		{
			"rounding up",
			`"3.499999999999 USD"`,
			`{"amount":3.49,"currency":"USD"}`,
		},
		{
			"from json amount as number",
			`{"amount":99,"currency":"PLN"}`,
			`{"amount":99.00,"currency":"PLN"}`,
		},
		{
			"from json amount as string",
			`{"amount":"88.12","currency":"CHF"}`,
			`{"amount":88.12,"currency":"CHF"}`,
		},
		{
			"big amount formating",
			`{"amount":13027806790,"currency":"PLN"}`,
			`{"amount":13027806790.00,"currency":"PLN"}`,
		},
		{
			"with empty amount",
			`{"currency":"PLN"}`,
			`money:not supported amount type`,
		},
		{
			"with empty currency",
			`{"amount":"1999","currency":""}`,
			`money:currency can not be empty`,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			var m Money
			if err := m.UnmarshalJSON([]byte(c.input)); err != nil {
				if err.Error() == c.output {
					return
				}
				t.Fatalf("no unmarshal error expected, got %s", err)
			}
			b, err := json.Marshal(m)
			if err != nil {
				t.Fatalf("no marshal error expected, got %s", err)
			}
			if string(b) != c.output {
				t.Fatalf("expected %s, got %s", c.output, string(b))
			}
		})
	}
}

func TestMoney_Scan(t *testing.T) {
	var m Money
	if err := m.Scan("Parking Lot: 1.86 USD"); err != nil {
		t.Fatal(err)
	}
	if *m.Amount != 186 {
		t.Fatal("expected 186, got", *m.Amount)
	}
	if m.Currency != "USD" {
		t.Fatal("expected USD, got", m.Currency)
	}
	if m.Label != "Parking Lot" {
		t.Fatal("expected parking lot, got", m.Label)
	}
}

func TestMoney_Add(t *testing.T) {
	m := MustMoney("It is a floating point stretch test: 0 USD")
	var err error
	for i := 1; i <= 1_000_000; i++ {
		if m, err = m.Add(MustMoney("0.01 USD")); err != nil {
			t.Fatal(err)
		}
	}
	if m.IsEmpty() {
		t.Fatalf("expected non-empty money")
	}
	if m.Float() != 10000 {
		t.Fatalf("expected 10000, got %g", m.Float())
	}
	if m.String() != "It is a floating point stretch test: 10000.00 USD" {
		t.Fatalf("expected 10000.00 USD, got %s", m.String())
	}
	if _, err = MustMoney("5.50 USD").Add(MustMoney("3 PLN")); err == nil {
		t.Fatalf("expected no errors when money with exact same currency are added")
	}
}

func TestMoney_Eq(t *testing.T) {
	a := MustMoney("32.76 PLN")
	b := MustMoney("32.76 PLN")
	c := MustMoney("32.76 USD")

	if !a.Eq(b) {
		t.Fatal()
	}
	if a.Eq(c) {
		t.Fatal()
	}
}
