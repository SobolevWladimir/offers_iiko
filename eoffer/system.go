package eoffer

import (
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/order_engine"
	"fmt"
	"time"
)

type System struct {
	Order *order_engine.Order
}

func (s System) GetValueForOfffer(field string, filters offerentity.OfferFilterValues) (interface{}, error) {
	switch field {
	case "date":
		return time.Now(), nil
	case "time":
		return time.Now(), nil
	case "request_by_url":
		return s.Order, nil
	}
	return nil, fmt.Errorf(" field %v not found", field)
}
