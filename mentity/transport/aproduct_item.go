package transport

import "offers_iiko/lib/base"

type ProductItemSityInfo struct {
	ID     int     `json:"id"`
	Price  float32 `json:"price"`
	Profit float32 `json:"profit"`
}
type ProductItem struct {
	ID       int                 `json:"id"`
	Name     string              `json:"name"`
	Weight   string              `json:"weight"`
	SityInfo ProductItemSityInfo `json:"sity_info"`
	Vendor1  base.StringInt      `json:"vendor1"`
	Vendor2  base.StringInt      `json:"vendor2"`
}
