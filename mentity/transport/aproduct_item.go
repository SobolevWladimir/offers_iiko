package transport

import "offers_iiko/lib/base"

type ProductItem struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Weight  string         `json:"weight"`
	Vendor1 base.StringInt `json:"vendor1"`
	Vendor2 base.StringInt `json:"vendor2"`
}
