package service

import (
	"altegra_offers/service/coupon"
	"altegra_offers/service/coupon_category"
	"altegra_offers/service/offers_engine"
)

func init() {
	//webhook_type
	/*	addService(webhook_type.New())
		//webhook
		addService(webhook.New())
		//address
		addService(city.New())
		addService(street.New())
		//status

		// token
		addService(token.New())
		// product
		addService(status.New())
		addService(modificationgroup.New())
		addService(product_categories.New())
		addService(product.New())
		addService(modifications.New())
		//pricelist
		addService(product_cost.New())
		addService(command_change.New())
		//order
		addService(client.New())
		addService(client_category.New())
		addService(client_category_link.New())
		addService(order_engine.New())
	*/
	// marketing
	addService(coupon_category.New())
	addService(coupon.New())
	addService(offers_engine.New())

	//sync module Должени быть всегда последним
	//	addService(sync.New())
}
