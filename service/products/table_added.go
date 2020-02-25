package products

import (
	"offers_iiko/lib/tablename"
	"offers_iiko/mentity/transport"
)

type ProductAdded struct {
	Id        int `db:"id"`
	ParentId  int `db:"parent_id"`
	ProductId int `db:"product_id"`
	Type      int `db:"type"`
	Quantity  int `db:"quantity"`
}
type ProductAddeds []ProductAdded

func FindAllProductsAddeds() (ProductAddeds, error) {
	result := ProductAddeds{}
	db := connect()
	err := db.Select(&result, `select * from `+tablename.ProductsAdded)
	return result, err
}
func (a *ProductAddeds) FindByParent(id int) ProductAddeds {
	result := ProductAddeds{}
	for _, prod := range *a {
		if prod.ParentId == id {
			result = append(result, prod)
		}
	}
	return result
}

func (a *ProductAddeds) FindByType(t int) ProductAddeds {
	result := ProductAddeds{}
	for _, prod := range *a {
		if prod.Type == t {
			result = append(result, prod)
		}
	}
	return result
}
func (a *ProductAddeds) GetByType(t int) ProductAdded {
	for _, prod := range *a {
		if prod.Type == t {
			return prod
		}
	}
	return ProductAdded{
		ProductId: 0,
	}
}
func (a *ProductAddeds) GetProductID() []int {
	result := []int{}
	for _, add := range *a {
		result = append(result, add.ProductId)
	}
	return result
}
func (a *ProductAdded) ToAProductAddedItem() (transport.ProductAddedItem, error) {

	result := transport.ProductAddedItem{}
	result.ID = a.Id
	result.Parent = a.ParentId
	result.Quantity = a.Quantity
	prod, err := storage_product.FindById(a.ProductId)
	if err != nil {
		return result, err
	}
	result.Product = prod.ToAProductItem()
	return result, nil
}
func (a *ProductAddeds) ToAProductAddedsItems() (transport.ProductAddedItems, error) {

	result := transport.ProductAddedItems{}
	for _, ad := range *a {
		item, err := ad.ToAProductAddedItem()
		if err != nil {
			return result, err
		}
		result = append(result, item)
	}
	return result, nil
}
