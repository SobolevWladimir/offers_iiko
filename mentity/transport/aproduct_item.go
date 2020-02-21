package transport

import (
	"offers_iiko/lib/base"
)

type ProductItemSityInfo struct {
	ID     int     `json:"id" `
	Price  float32 `json:"price"`
	Profit float32 `json:"profit"`
}
type AProductItem struct {
	ID                    int                 `json:"id"`
	Name                  string              `json:"name"`
	Weight                string              `json:"weight"`
	SityInfo              ProductItemSityInfo `json:"sity_info"`
	Vendor1               base.StringInt      `json:"vendor1"`
	Vendor2               base.StringInt      `json:"vendor2"`
	Type                  int                 `json:"-"`
	Comment               string              `json:"comment"`
	Alias                 string              `json:"alias"`
	Number                string              `json:"number"`
	Caloric               float32             `json:"caloric"`
	Pfc                   string              `json:"pfs"`
	Composition           string              `json:"composition"`
	Image                 string              `json:"image"`
	FullImage             string              `json:"full_image"`
	New                   bool                `json:"new"`
	Hot                   bool                `json:"hot"`
	Hit                   bool                `json:"hit"`
	NotDeliverySeparately bool                `json:"not_delivery_separately"`
	Volume                float32             `json:"volume"`
	Article               string              `json:"article"`
}
