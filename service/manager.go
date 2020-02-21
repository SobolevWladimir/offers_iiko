package service

import (
	"offers_iiko/service/iiko"
	"offers_iiko/service/products"
	site_setting "offers_iiko/service/setting"
	"offers_iiko/service/webhook"
)

func init() {
	addService(webhook.New())
	addService(site_setting.New())
	addService(iiko.New())
	addService(products.New())
}
