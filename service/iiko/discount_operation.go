package iiko

import (
	"offers_iiko/mentity/offerentity"
	"offers_iiko/mentity/transport"
)

type OperationCode int

const (
	OperationCodeOrderFixedSumDiscount = 0
	OperationCodeOFixedSumDiscount     = 1
	OperationCodeFreeProduct           = 2
)

type DiscountOperation struct {
	Code        OperationCode `json:"code"` //Код типа операции
	OrderItemId string        `json:"orderItemId"`
	ProductCode string        `json:"productCode"`
	ProductName string        `json:"productName"`
	DiscountSum float32       `json:"discountSum"`
	Comment     string        `json:"comment"`
}
type DiscountOperations []DiscountOperation

func (ds *DiscountOperations) GetActons(order transport.IOrderRequest) (offerentity.Actions, error) {
	result := offerentity.Actions{}
	for _, d := range *ds {
		target, err := order.Order.Items.GetSiteIdByOrderItemId(d.OrderItemId)
		if err != nil {
			return result, err
		}
		action := offerentity.ActionSpecialDiscount{
			Max:         1,
			SaleType:    1,
			Target:      target,
			Value:       d.DiscountSum,
			OrderItemId: d.OrderItemId,
		}
		result = append(result, offerentity.Action{
			Type: offerentity.TypeSpecialDiscount,
			Data: action,
		})
	}

	return result, nil
}
