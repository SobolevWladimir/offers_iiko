package iiko

import "errors"

type Product struct {
	Id                string `json:"id"`
	Code              string `json:"code"`
	Name              string `json:"name"`
	ProductCategoryId string `json:"productCategoryId"`
}
type Products []Product

func (ps *Products) FindProductByCode(code string) (Product, error) {
	for _, p := range *ps {
		if p.Code == code {
			return p, nil
		}
	}
	return Product{}, errors.New("Не найден продукт с кодом:" + code)
}
