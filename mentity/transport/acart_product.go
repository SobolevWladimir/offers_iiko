package transport

import "offers_iiko/lib/base"

type ACartProducts []ACartProduct
type ACartProduct struct {
	Product     AProduct     `json:"product"`
	Category    string       `json:"category"`
	OrderItemId string       `json:"orderItemId"`
	Quantity    int          `json:"quantity"`
	Options     ACartOptions `json:"options"`
}

func (a *ACartProducts) ToOrderItems(nm NomenclatureInterface) IOrderItems {
	result := IOrderItems{}
	for _, item := range *a {
		result = append(result, item.ToIOrderItem(nm))
	}
	return result
}
func (a *ACartProduct) ToIOrderItem(nm NomenclatureInterface) IOrderItem {
	result := IOrderItem{}
	if len(a.OrderItemId) > 0 {
		result.OrderItemId = a.OrderItemId
	} else {
		result.OrderItemId = base.UUID()
	}
	product_id, err := nm.FindProductIdByCode(string(a.Product.Vendor1))
	result.ID = product_id
	if err != nil {
		result.ID = base.UUID()
	}
	result.Code = string(a.Product.Vendor1)
	result.Amount = float32(a.Quantity)
	result.Sum = a.Product.SityInfo.Price * result.Amount
	result.SiteId = a.Product.ID
	cat, _ := nm.FindCategoryNameByProductCode(string(a.Product.Vendor1))
	result.Cagegory = cat
	return result
}
func (a *ACartProduct) GetIOrderItemModifiers() IOrderItemModifiers {
	result := IOrderItemModifiers{}
	if a.Options.Added.DoubleMeat != 0 {
		result = append(result, IOrderItemModifier{
			ID:     base.UUID(),
			Amount: 1,
		})
	}
	return result
}
