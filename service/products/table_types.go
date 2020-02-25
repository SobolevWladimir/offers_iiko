package products

import "offers_iiko/lib/tablename"

type ProductType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type ProductTypes []ProductType

func FindAllProductTypes() (ProductTypes, error) {
	result := ProductTypes{}
	db := connect()
	err := db.Select(&result, `select * from `+tablename.ProductTypes)
	return result, err
}
func (t *ProductTypes) GetById(id int) ProductType {
	for _, typ := range *t {
		if typ.Id == id {
			return typ
		}
	}
	return ProductType{}

}
