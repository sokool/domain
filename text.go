package domain

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Text struct {
	string
	min, max uint
}

func NewText(s string, min, max uint) (_ Text, err error) {
	t := Text{
		string: s,
		min:    min,
		max:    max,
	}

	return t, t.valid(s)
}

func MustText(s string) Text {
	t, err := NewText(s, 0, 64*1024)
	if err != nil {
		panic(err)
	}
	return t
}

func (t Text) Set(s string) (Text, error) {
	if err := t.valid(s); err != nil {
		return t, err
	}

	t.string = s
	return t, nil
}

func (t Text) IsZero() bool {
	return t.string == ""
}

func (t *Text) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, &t.string)
}

func (t Text) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.string)
}

func (t Text) UUID() uuid.UUID {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(t.string))
}

func (t Text) Length() int { return len(t.string) }

func (t Text) LengthBetween(from, to int) bool { n := t.Length(); return n >= from && n <= to }

func (t Text) Append(ss ...fmt.Stringer) Text {
	var vv []string
	for i := range ss {
		if ss[i].String() == "" {
			continue
		}

		vv = append(vv, ss[i].String())
	}

	if len(vv) == 0 {
		return t
	}

	t.string = fmt.Sprintf("%s %s", t, strings.Join(vv, " "))
	return t
}

func (t Text) Word(n int) Text {
	if w := t.split(" "); !t.IsZero() && len(w) > n {
		t.string = w[n]
		return t
	}

	return Text{}
}

func (t Text) Is(s string) bool {
	return t.string == s
}

func (t Text) Number() (float64, error) {
	return strconv.ParseFloat(t.String(), 64)
}

func (t Text) Words() int { return len(t.split(" ")) }

func (t Text) Contains(s string) bool { return strings.Contains(t.string, s) }

func (t Text) Print(w ...io.Writer) {
	fmt.Fprintln(os.Stdout, t)
}

func (t *Text) split(sep string) []string {
	return strings.Split(t.String(), sep)
}

func (t Text) String() string {
	return t.string
}

func (t Text) valid(s string) error {
	if t.min > t.max {
		return Errorf("max can not be greater than min")
	}
	if n := uint(len(s)); n < t.min || n > t.max {
		return Errorf("`%s` has invalid length, required between %d-%d characters", s, t.min, t.max)
	}
	return nil
}
