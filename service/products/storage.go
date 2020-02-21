package products

var storage_product ProductItems

func UpdataStorage() error {
	prods, err := FindAllProducts()
	if err != nil {
		return err
	}
	storage_product = prods
	return nil
}
