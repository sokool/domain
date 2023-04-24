package domain

//import (
//	"fmt"
//	"strconv"
//)
//
//type Percent float64
//
//func NewPercent(p float64) Percent {
//	return Percent(round(p, 6))
//}
//
//func ParsePercent(s string) (Percent, error) {
//	var p Percent
//
//	return p, p.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, s)))
//}
//
//func (p Percent) Between(from, to Percent) bool {
//	return !(p > to || p < from)
//}
//
//func (p Percent) Normalize() float64 {
//	return float64(p) / 100
//}
//
////func (p Percent) Of(c Money) Money {
////	return c.Multi(float64(p / 100))
////}
//
//func (p Percent) String() string {
//	return fmt.Sprintf("%.2f%%", p)
//}
//
//func (p Percent) MarshalJSON() ([]byte, error) {
//	return []byte(fmt.Sprintf(`"%s"`, strconv.FormatFloat(float64(p), 'f', -1, 64))), nil
//}
//
//func (p *Percent) UnmarshalJSON(b []byte) error {
//	n := len(b)
//	if n <= 2 {
//		return nil
//	}
//
//	f, err := strconv.ParseFloat(string(b[1:n-1]), 32)
//	if err != nil {
//		return err
//	}
//
//	*p = NewPercent(f)
//	return nil
//}
//
//type Tax struct{ Percent }
//
//func NewTax(percent float64) (Tax, error) {
//	p := NewPercent(percent)
//	if !p.Between(0, 100) {
//		return Tax{p}, Errorf("tax: expected 0%% - 100%%, got %s", p.String())
//	}
//	return Tax{p}, nil
//}
