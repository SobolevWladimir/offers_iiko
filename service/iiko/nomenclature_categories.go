package iiko

import "errors"

type ProductCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ProductCategorys []ProductCategory

func (ps *ProductCategorys) FindByID(id string) (ProductCategory, error) {
	for _, p := range *ps {
		if p.ID == id {
			return p, nil
		}
	}
	return ProductCategory{}, errors.New("Не могу найти категорию. id:" + id)
}
