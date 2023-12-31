package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

var ErrorProdNotFound = fmt.Errorf("product not found")

func GetProducts() Products {

	return productList
}
func nextProdId() int {
	l := productList[len(productList)-1]
	return l.ID + 1
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)

}
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	mathces := re.FindAllString(fl.Field().String(), -1)
	if len(mathces) != 1 {
		return false
	}
	return true
}
func AddProducts(prod *Product) {
	prod.ID = nextProdId()
	productList = append(productList, prod)
}

func UpdateProducts(id int, prod *Product) error {
	_, pos, err := FindProduct(id)
	if err != nil {
		return err
	}
	prod.ID = id
	productList[pos] = prod
	return nil
}

func FindProduct(id int) (*Product, int, error) {
	for i, prod := range productList {
		if prod.ID == id {
			return prod, i, nil
		}
	}
	return nil, -1, ErrorProdNotFound

}
func (prods *Products) ToJSON(res io.Writer) error {
	e := json.NewEncoder(res)
	return e.Encode(prods)
}
func (prod *Product) FromJSON(req io.Reader) error {
	d := json.NewDecoder(req)
	return d.Decode(prod)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk",
		Price:       1.99,
		SKU:         "fgj789",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
