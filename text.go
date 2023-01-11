package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Text string

func NewText(s string) Text {
	return Text(s)
}

func (t Text) Length() int { return len(t) }

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

	return Text(fmt.Sprintf("%s %s", t, strings.Join(vv, " ")))
}

func (t Text) IsZero() bool { return strings.TrimSpace(string(t)) == "" }

func (t Text) Word(n int) Text {
	if w := t.split(" "); !t.IsZero() && len(w) > n {
		return Text(w[n])
	}

	return ""
}

func (t Text) Number() (float64, error) {
	return strconv.ParseFloat(t.String(), 64)
}

func (t Text) Words() int { return len(t.split(" ")) }

func (t Text) String() string { return string(t) }

func (t Text) Contains(s string) bool { return strings.Contains(string(t), s) }

func (t Text) Print(w ...io.Writer) {
	fmt.Fprintln(os.Stdout, t)
}

func (t *Text) split(sep string) []string {
	return strings.Split(t.String(), sep)
}

func (t Text) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}

	return json.Marshal(t.String())
}

func (t *Text) UnmarshalJSON(bb []byte) error {
	n := len(bb)
	if n <= 2 || bytes.Equal(bb, []byte(`null`)) {
		return nil
	}

	var s string
	if err := json.Unmarshal(bb, &s); err != nil {
		return err
	}

	*t = Text(s)
	return nil
}

func (t *Text) Scan(v interface{}) error {
	switch s := v.(type) {
	case []byte:
		*t = Text(s)
	default:
		return Err("text", "incompatible type for Body")
	}

	return nil
}
