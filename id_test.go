package domain_test

import (
	"fmt"
	"testing"

	"github.com/sokool/domain"
)

func TestNewID(t *testing.T) {
	//_, err := domain.NewHash[string]()
	//if err == nil {
	//	t.Fatal(err)
	//}
	var a domain.ID[int]

	x, _ := domain.NewID[string]()
	y, _ := domain.NewID[Product]()

	fmt.Println(x, x.IsEmpty())
	fmt.Println(y, y.IsEmpty())
	fmt.Println(a, a.IsEmpty())

	//fmt.Println(id)
	//if err := domain.NewID(&id, "86b7cdbb-1c03-409e-a6e9-65501ff97036"); err != nil {
	//	t.Fatal(err)
	//}

	//fmt.Println(id)
}

type Product struct {
	ID domain.ID[Product]
}
