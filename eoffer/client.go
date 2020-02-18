package eoffer

import (
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/client"
	"altegra_offers/service/client_category_link"
	"altegra_offers/service/order_engine"
	"fmt"
)

type Client string

func (c Client) GetValueForOfffer(field string, filters offerentity.OfferFilterValues) (interface{}, error) {
	if len(c) == 0 {
		return nil, nil
	}
	switch field {
	case "phone":
		return c, nil
	case "bonuses":
		{
			client, _ := client.FindOneByPhone(string(c))
			return client.Bonuses, nil
		}
	case "categories":
		{
			catgs, _ := client_category_link.FindAllCategoriesByClient(string(c))
			return catgs, nil
		}
	case "birth_date":
		{
			client, _ := client.FindOneByPhone(string(c))
			return client.BirthDate, nil
		}
	case "sex":
		{
			client, _ := client.FindOneByPhone(string(c))
			return client.Sex, nil
		}
	case "orders_product":
		return order_engine.GetClietnProductCount(string(c), filters)
	case "category":
		return client_category_link.FindAllCategoriesByClient(string(c))
	case "last_activity":
		return order_engine.GetLastActivityClient(string(c), filters)
	case "order_count":
		return order_engine.GetCountOrdersClient(string(c), filters)

	}
	return nil, fmt.Errorf("field %v not found", field)
}
