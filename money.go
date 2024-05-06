package domain

import (
	"database/sql/driver"
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Money ...
type Money struct {
	Amount   *amount  `json:"amount"`
	Currency Currency `json:"currency"`
	Label    string   `json:"label,omitempty"`
}

type money interface {
	float64 | float32 | int | int32 | int64
}

func NewMoney[M money](value M, currency string, label ...string) Money {
	var f amount
	switch v := any(value).(type) {
	case float64:
		f = amount(v * math.Pow10(2))
	case int:
		f = amount(v * 100)
	case int32:
		f = amount(v * 100)
	}
	return Money{Amount: &f, Currency: Currency(currency), Label: strings.Join(label, " ")}
}

func ParseMoney(s string, args ...any) (Money, error) {
	if len(args) != 0 {
		s = fmt.Sprintf(s, args...)
	}
	p := reMoney.FindStringSubmatch(s)
	if len(p) != 4 {
		return Money{}, ErrMoney.New("invalid '%s' string format", s)
	}
	f, err := strconv.ParseFloat(p[2], 64)
	if err != nil {
		return Money{}, ErrMoney.New("invalid '%s' string format", s)
	}
	return NewMoney(f, strings.TrimSpace(p[3]), strings.TrimSpace(strings.ReplaceAll(p[1], ":", ""))), nil
}

func MustMoney(s string, args ...any) Money {
	m, err := ParseMoney(s, args...)
	if err != nil {
		panic(err)
	}
	return m
}

func (m Money) Add(n Money) (Money, error) {
	if !m.isOK(n) {
		return Money{}, ErrMoney.New("currencies are different or Money is empty value")
	}
	s := *m.Amount + *n.Amount
	return Money{Amount: &s, Currency: m.Currency, Label: m.Label}, nil
}

func (m Money) Multi(f float64) Money {
	if m.IsEmpty() {
		return Money{}
	}
	s := amount(float64(*m.Amount) * f)
	return Money{Amount: &s, Currency: m.Currency, Label: m.Label}
}

func (m Money) IsEmpty() bool {
	return m.Amount == nil
}

func (m Money) Le(n Money) bool {
	return m.isOK(n) && *m.Amount <= *n.Amount
}

func (m Money) Lt(n Money) bool {
	return m.isOK(n) && *m.Amount < *n.Amount
}

func (m Money) Eq(n Money) bool {
	return m.isOK(n) && *m.Amount == *n.Amount
}

func (m Money) Gt(n Money) bool {
	return m.isOK(n) && *m.Amount > *n.Amount
}

func (m Money) Ge(n Money) bool {
	return m.isOK(n) && *m.Amount >= *n.Amount
}

func (m Money) Convert(to Currency, c Converter) (Money, error) {
	if m.Currency == to {
		return m, nil
	}
	n, err := c.Convert(m, to)
	if err != nil {
		return Money{}, ErrMoney.New("convert: %w", err)
	}
	if n.Label == "" {
		n.Label = m.Label
	}
	return n, nil
}

func (m Money) String() string {
	if m.IsEmpty() {
		return ""
	}
	if m.Label != "" {
		return fmt.Sprintf("%s: %.2f %s", m.Label, m.Float(), m.Currency)
	}
	return strings.TrimSpace(fmt.Sprintf("%.2f %s", m.Float(), m.Currency))
}

func (m Money) GoString() string {
	b, _ := json.MarshalIndent(m, "", "\t")
	return fmt.Sprintf("%T\n%s\n", m, b)
}

func (m Money) MarshalJSON() (b []byte, err error) {
	if m.IsEmpty() {
		return []byte("null"), nil
	}
	type void Money
	if b, err = json.Marshal((*void)(&m)); err != nil {
		return nil, ErrMoney.Wrap(err)
	}
	return
}

func (m *Money) UnmarshalJSON(b []byte) error {
	var err error
	switch n := len(b); {
	case n == 0 || string(b) == `null`:
		return nil
	case n > 2 && b[0] == '"':
		*m, err = ParseMoney(string(b[1 : n-1]))
		return err
	case n > 2 && b[0] == '{':
		var j struct {
			Amount          any
			Currency, Label string
		}
		if err = json.Unmarshal(b, &j); err != nil {
			return ErrMoney.Wrap(err)
		}
		if j.Currency == "" {
			return ErrMoney.New("currency can not be empty")
		}
		switch a := j.Amount.(type) {
		case string:
			*m, err = ParseMoney("%s:%v%s", j.Label, a, j.Currency)
		case float64:
			*m, err = ParseMoney("%s:%f%s", j.Label, a, j.Currency)
		default:
			return ErrMoney.New("not supported amount type")
		}
		return nil
	default:
		*m, err = ParseMoney(string(b))
		return err
	}
}

func (m *Money) Scan(src any) error {
	switch v := src.(type) {
	case string:
		if err := m.UnmarshalJSON([]byte(v)); err != nil {
			return ErrMoney.Wrap(err)
		}
		return nil
	}
	return ErrMoney.New("scan: %T type not suported", src)
}

func (m Money) Value() (driver.Value, error) {
	return m.Amount, nil
}

func (m Money) isOK(n Money) bool {
	return !(m.IsEmpty() || n.IsEmpty() || m.Currency != n.Currency)
}

func (m Money) IsPositive(equal ...bool) bool {
	return *m.Amount > 0
}

func (m Money) IsNegative() bool {
	return *m.Amount < 0
}

func (m Money) Render(name ...string) Meta {
	return Meta{
		"amount":   m.Amount,
		"currency": m.Currency,
	}
}

// Float round Money to float64
//
// This code just do rounding to support many common use cases on Money, however it is
// not precise for large money numbers and operations. In case of good precision please
// use e.g. https://github.com/shopspring/decimal
func (m Money) Float() float64 {
	if m.Amount == nil {
		return math.NaN()
	}
	return float64(*m.Amount) / 100
}

type amount int64

func (a *amount) Float() float64 {
	if a == nil {
		return math.NaN()
	}
	return float64(*a) / 100
}

func (a *amount) MarshalJSON() ([]byte, error) {
	return []byte(a.String()), nil
}

func (a *amount) UnmarshalJSON(b []byte) error {
	f, err := strconv.ParseFloat(string(b), 64)
	*a = amount(f * 100)
	return err
}

func (a *amount) String() string {
	if a == nil {
		return ""
	}
	return fmt.Sprintf("%.2f", a.Float())
}

var (
	ErrMoney = Errorf("money")
	reMoney  = regexp.MustCompile(`^(.*?:\s?)?(-?\d*\.?\d+)(\s?[A-Z]{3})$`)
)

var (
	EUR Currency = "EUR"
	USD Currency = "USD"
)

type Currency string

func (c Currency) String() string {
	return strings.ToUpper(string(c))
}

type Converter interface {
	Convert(Money, Currency) (Money, error)
}

type converter bool

var FakeConverter = converter(true)

func (c converter) Convert(from Money, to Currency) (Money, error) {
	var r float64 = 1
	if to == "PLN" {
		switch from.Currency {
		case "USD":
			r = 3.98
		}

		m := from.Float() * r
		return NewMoney(m, string(to)), nil
	}

	if to != USD {
		return Money{}, fmt.Errorf("%w:converter: currency %s not supported", ErrMoney, to)
	}
	switch from.Currency {
	case "NZD":
		r = 0.62
	case "EUR":
		r = 1.0889
	case "AUD":
		r = 0.66
	case "SGD":
		r = 0.75
	case "PLN":
		r = 0.2535
	case "CHF":
		r = 1.11
	default:
		return Money{}, fmt.Errorf("%w:converter: %s not supported", ErrMoney, from.Currency)
	}
	m := from.Float() * r
	return NewMoney(m, string(to)), nil
}
