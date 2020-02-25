package products

var storage_product ProductItems
var storage_product_added ProductAddeds
var storage_product_types ProductTypes

func UpdataStorage() error {
	//prods
	prods, err := FindAllProducts()
	if err != nil {
		return err
	}
	storage_product = prods

	//added
	adeds, err := FindAllProductsAddeds()
	if err != nil {
		return err
	}
	storage_product_added = adeds

	//types
	types, err := FindAllProductTypes()
	if err != nil {
		return err
	}
	storage_product_types = types

	return nil
}
