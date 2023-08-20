package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk",
		Price:       1.99,
		SKU:         "fgj789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
