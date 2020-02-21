package transport

import "fmt"

type ProductAddedItem struct {
	ID      int          `json:"id"`
	Product AProductItem `json:"product"`
}
type ProductAddedItems []ProductAddedItem

type ProductAdded struct {
	Meat      ProductAddedItem  `json:"meat"`
	Vegetable ProductAddedItem  `json:"vegetable"`
	Noodle    ProductAddedItems `json:"noodle"`
	Toping    ProductAddedItems `json:"toping"`
	Souse     ProductAddedItems `json:"souse"`
}

func (p *ProductAddedItems) GetVendor1ByProductId(id int) (string, error) {
	for _, item := range *p {
		if item.Product.ID == id {
			return string(item.Product.Vendor1), nil
		}
	}
	return "", fmt.Errorf("Не могу найти продукт(%v) в  added", id)
}
