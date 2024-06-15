package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Price     float32 `json:"price"`
	SKU       string  `json:"sku"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
}
type ProductList []*Product

func (p *ProductList) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func GetProducts() ProductList {
	return Products
}

func AddProducts(p *Product) {
	p.ID = getNextId()
	Products = append(Products, p)
}
func getNextId() int {
	lp := Products[len(Products)-1]
	return lp.ID + 1
}

func UpdateProducts(id int, p *Product) error {
	_, pos, err := FindProductById(id)
	if err != nil {
		return err
	}
	p.ID = id
	Products[pos] = p
	return nil
}

var ErrProductNotFound = errors.New("product not found")

func FindProductById(id int) (*Product, int, error) {
	for i, product := range Products {
		if product.ID == id {
			return product, i, nil
		}
	}
	return nil, 0, ErrProductNotFound
}

var Products = []*Product{
	&Product{
		ID:        1,
		Name:      "satya",
		Desc:      "Siemens",
		Price:     50.1,
		SKU:       "Z004KZEP",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Saurabh",
		Desc:      "Saurabh",
		Price:     50.2,
		SKU:       "Z004K909",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
