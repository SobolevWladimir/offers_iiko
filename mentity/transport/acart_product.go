package transport

import "offers_iiko/lib/base"

type ACartProducts []ACartProduct
type ACartProduct struct {
	Product  AProduct     `json:"product"`
	Quantity int          `json:"quantity"`
	Options  ACartOptions `json:"options"`
	Mark     string       `json:"mark"` //  отправить если  небходимо назначить свой  uuid
}

func (a *ACartProducts) ToOrderItems() IOrderItems {
	result := IOrderItems{}
	for _, item := range *a {
		result = append(result, item.ToIOrderItem())
	}
	return result
}
func (a *ACartProduct) ToIOrderItem() IOrderItem {
	result := IOrderItem{}
	if len(a.Mark) > 0 {
		result.ID = a.Mark
	} else {
		result.ID = base.UUID()
	}
	result.Code = string(a.Product.Vendor1)
	result.Amount = float32(a.Quantity)
	result.Sum = a.Product.SityInfo.Price * result.Amount
	result.SiteId = a.Product.ID
	//result.Modifiers = a.GetIOrderItemModifiers()
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
