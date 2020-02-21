package products

import (
	"offers_iiko/mentity/transport"
)

type TableProduct struct {
}

func (t *TableProduct) GetProductByCode(code string) (transport.AProductItem, error) {
	prod, err := FindProductByVendor1(code)
	return prod.ToAProductItem(), err
}
func NewTableProduct() *TableProduct {
	return new(TableProduct)
}
