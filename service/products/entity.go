package products

import (
	"offers_iiko/mentity/transport"
)

const (
	ProductTypeNoodle    = 3
	ProductTypeMeat      = 4
	ProductTypeComposite = 6 //	 composite
	ProductTypeVagetable = 7 //
	ProductTypeToping    = 8
	ProductTypeSouse     = 10
)

type TableProduct struct {
}

func (t *TableProduct) GetProductByCode(code string) (transport.AProduct, error) {

	result := transport.AProduct{}

	prod, err := FindProductByVendor1(code)
	result.AProductItem = prod.ToAProductItem()
	added := storage_product_added.FindByParent(prod.ID)
	//composite
	added_composite := added.FindByType(ProductTypeComposite)
	composite, err := added_composite.ToAProductAddedsItems()
	if err != nil {
		return result, err
	}
	result.Added.Composite = composite
	//meat
	added_meat := added.GetByType(ProductTypeMeat)
	if added_meat.ProductId != 0 {
		meat, err := added_meat.ToAProductAddedItem()
		if err != nil {
			return result, err
		}
		result.Added.Meat = meat

	}
	added_vegetable := added.GetByType(ProductTypeVagetable)
	if added_vegetable.ProductId != 0 {
		vegetable, err := added_vegetable.ToAProductAddedItem()
		if err != nil {
			return result, err
		}
		result.Added.Vegetable = vegetable
	}
	added_noodle := added.FindByType(ProductTypeNoodle)
	noodle, err := added_noodle.ToAProductAddedsItems()
	if err != nil {
		return result, err
	}
	result.Added.Noodle = noodle

	added_toping := added.FindByType(ProductTypeToping)
	toping, err := added_toping.ToAProductAddedsItems()
	if err != nil {
		return result, err
	}
	result.Added.Toping = toping

	added_souse := added.FindByType(ProductTypeSouse)
	souse, err := added_souse.ToAProductAddedsItems()
	if err != nil {
		return result, err
	}
	result.Added.Souse = souse
	return result, err
}
func NewTableProduct() *TableProduct {
	return new(TableProduct)
}
