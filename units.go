package domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	Kilo float64 = 1e+3
	Mega float64 = 1e+6
	Giga float64 = 1e+9
	Tera float64 = 1e+12

	Kbps = Throughput(Kilo)
	Mbps = Throughput(Mega)
	Gbps = Throughput(Giga)
	Tbps = Throughput(Tera)

	KHz = Frequency(Kilo)
	MHz = Frequency(Mega)
	GHz = Frequency(Giga)
	THz = Frequency(Tera)

	KB = Capacity(Kilo)
	MB = Capacity(Mega)
	GB = Capacity(Giga)
	TB = Capacity(Tera)

	KW = Power(Kilo)
	MW = Power(Mega)
	GW = Power(Giga)
	TW = Power(Tera)
)

type (
	Throughput = Unit[bps]
	Capacity   = Unit[bytez]
	Frequency  = Unit[hertzs]
	Power      = Unit[watts]
)

func ParseThroughput(s string) (Throughput, error) { return ParseUnit[bps](s) }

func ParseCapacity(s string) (Capacity, error) { return ParseUnit[bytez](s) }

func ParseFrequency(s string) (Frequency, error) { return ParseUnit[hertzs](s) }

func ParsePower(s string) (Power, error) { return ParseUnit[watts](s) }

type watts struct{}

func (w watts) Format(unit string) string { return unit + "W" }

type hertzs struct{}

func (h hertzs) Format(unit string) string { return unit + "Hz" }

type bps struct{}

func (b bps) Format(unit string) string { return unit + "bps" }

type bytez struct{}

func (b bytez) Format(unit string) string { return strings.ToUpper(unit) + "B" }

type Formatter interface {
	Format(unit string) string
}

type Unit[T Formatter] float64

// ParseUnit
// todo
//   - support negative values like -45.88MB
//   - unit can be based on int64 instead float64
func ParseUnit[T Formatter](s string) (Unit[T], error) {
	p := numUnit.FindStringSubmatch(strings.ReplaceAll(s, " ", ""))
	if len(p) != 3 {
		return -1, fmt.Errorf("%w: invalid format, use number followed by string unit name, ie 45.99MB or 56kHz", ErrUnit)
	}
	f, err := strconv.ParseFloat(p[1], 64)
	if err != nil {
		return -1, fmt.Errorf("%w: not a number", ErrUnit)
	}
	if p[2] == `` {
		return Unit[T](f), nil
	}
	var t T
	switch u := string(p[2][0]); {
	case u == "k":
		f, s = f*Kilo, t.Format(u)
	case u == "M":
		f, s = f*Mega, t.Format(u)
	case u == "G":
		f, s = f*Giga, t.Format(u)
	case u == "T":
		f, s = f*Tera, t.Format(u)
	default:
		s = t.Format("")
	}
	if s != p[2] {
		return -1, fmt.Errorf("%w: expected %s, got %s", ErrUnit, s, p[2])
	}
	return Unit[T](f), nil
}

func (u Unit[T]) Values(format ...Unit[T]) (float64, string) {
	if len(format) == 1 {
		_, s := format[0].Values()
		return u.Float() / format[0].Float(), s
	}

	var t T
	var f string
	var v = float64(u)
	switch {
	case v < Kilo:
		f = t.Format("")
	case v >= Kilo && v < Mega:
		v, f = v/Kilo, t.Format("k")
	case v >= Mega && v < Giga:
		v, f = v/Mega, t.Format("M")
	case v >= Giga && v < Tera:
		v, f = v/Giga, t.Format("G")
	default:
		v, f = v/Tera, t.Format("T")
	}

	return v, f
}

func (u Unit[T]) Kilo() float64 { return float64(u) / Kilo }

func (u Unit[T]) Mega() float64 { return float64(u) / Mega }

func (u Unit[T]) Giga() float64 { return float64(u) / Giga }

func (u Unit[T]) Tera() float64 { return float64(u) / Tera }

func (u Unit[T]) Float() float64 { return float64(u) }

func (u Unit[T]) GoString() string {
	var f, s = u.Values()
	// when the float has more than 2 decimal places represent it as string with ~ as a
	// prefix, ie ~34.34MB
	if a := strings.SplitN(fmt.Sprintf("%v", f), ".", 2); len(a) == 2 && len(a[1]) > 2 {
		return fmt.Sprintf("~%.2f%s", f, s)
	}
	return fmt.Sprintf("%#v%s", f, s)
}

func (u Unit[T]) String() string { f, s := u.Values(); return fmt.Sprintf("%#v%s", f, s) }

func (u *Unit[T]) UnmarshalText(b []byte) (err error) { *u, err = ParseUnit[T](string(b)); return }

func (u Unit[T]) MarshalText() ([]byte, error) { return []byte(u.String()), nil }

func (u Unit[T]) Meta(format Unit[T]) Meta {
	v, s := u.Values(format)
	return Meta{
		"value": v,
		"unit":  s,
	}
}

var numUnit = regexp.MustCompile(`^([0-9.,]+)([A-z]+)$`)
var ErrUnit = fmt.Errorf("unit")
