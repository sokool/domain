package domain

//func TestPercent(t *testing.T) {
//	if NewPercent(3) != 3 {
//		t.Fatalf("expected 3%%")
//	}
//
//	if c := NewPercent(10).Of(NewMoney(1.50)); c != 0.15 {
//		t.Fatalf("expected $0.15, got %s", c)
//	}
//
//	if s := NewPercent(5); s != 5.00 {
//		t.Fatalf("expected 5.00%%, got %s", s)
//	}
//}

//func TestPercent_JSON(t *testing.T) {
//	b, err := NewPercent(31.508181).MarshalJSON()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if s := string(b); s != `"31.508181"` {
//		t.Fatalf(`expected "31.508181", got %s`, s)
//	}
//
//	var p Percent
//	if err = p.UnmarshalJSON([]byte(`"15.7578"`)); err != nil {
//		t.Fatal(err)
//	}
//
//	if p != 15.7578 {
//		t.Fatalf(`expected "15.7578", got %s`, p)
//	}
//
//	if err = p.UnmarshalJSON([]byte(`""`)); err != nil {
//		t.Fatal(err)
//	}
//
//}
//
//func TestTax(t *testing.T) {
//	if _, err := NewTax(-1); err == nil {
//		t.Fatal()
//	}
//
//	if _, err := NewTax(0); err != nil {
//		t.Fatal(err)
//	}
//
//	if _, err := NewTax(100); err != nil {
//		t.Fatal(err)
//	}
//
//	if _, err := NewTax(101); err == nil {
//		t.Fatal()
//	}
//}
