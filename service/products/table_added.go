package products

import "offers_iiko/lib/tablename"

type ProductAdded struct {
	Id        int `db:"id"`
	ParentId  int `db:"parent_id"`
	ProductId int `db:"product_id"`
	Type      int `db:"type"`
}
type ProductAddeds []ProductAdded

func FindAllProductsAddeds() (ProductAddeds, error) {
	result := ProductAddeds{}
	db := connect()
	err := db.Select(&result, `select * from `+tablename.ProductsAdded)
	return result, err
}
