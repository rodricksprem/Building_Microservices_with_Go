package data

import "testing"

func TestChecksValidator(t *testing.T) {
	p := &Product{
		Name:        "Latte",
		Description: "Milky coffee",
		Price:       2.45,
		SKU:         "abc-efgikl",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
