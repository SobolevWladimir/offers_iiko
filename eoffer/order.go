package eoffer

import (
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/order_engine"
	"fmt"
)

type Order order_engine.Order

func (o Order) GetValueForOfffer(field string, filters offerentity.OfferFilterValues) (interface{}, error) {
	switch field {
	case "id":
		return o.Id, nil
	case "point":
		return o.Point, nil
	case "status":
		return o.Status, nil
	case "delivery":
		return o.Delivery, nil
	case "cook_in_date":
		return o.CookInDate, nil
	case "cook_in_time":
		return o.CookInTime, nil
	case "paid":
		return o.Paid, nil
	case "pre_amount":
		return o.PreAmount, nil
	case "amount":
		return o.Amount, nil
	case "person":
		return o.Person, nil
	case "address":
		return o.Address, nil
	case "markers":
		return o.Markers, nil
	case "payment":
		return o.Payments, nil
	case "products":
		return o.Products.GetIds(), nil
	case "coupon":
		return o.Coupon, nil
	}
	return nil, fmt.Errorf(" field %v not found", field)
}
