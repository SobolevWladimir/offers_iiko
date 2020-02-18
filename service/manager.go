package service

import "offers_iiko/service/webhook"

func init() {
	addService(webhook.New())
}
