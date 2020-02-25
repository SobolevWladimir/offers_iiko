package iiko

import "offers_iiko/mentity/transport"

const (
	BizScheme = "https"
	BizHost   = "iiko.biz:9900"
)

type TableProduct interface {
	GetProductByCode(code string) (transport.AProduct, error)
}
