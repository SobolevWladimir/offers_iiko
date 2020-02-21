package products

import (
	"offers_iiko/lib/tablename"
)

func FindProductByVendor1(vendor string) (ProductItem, error) {
	result := ProductItem{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM `+tablename.Product+`  where vendor1=?`, vendor)
	return result, err
}
func FindAllProducts() (ProductItems, error) {
	result := ProductItems{}
	db := connect()
	err := db.Select(&result, `select * from `+tablename.Product)
	return result, err
}
