package products

import (
	"offers_iiko/lib/base"
	"offers_iiko/mentity/transport"

	"gopkg.in/guregu/null.v3"
)

type ProductItem struct {
	ID                    int         `db:"id"`
	Type                  int         `db:"type"`
	Name                  string      `db:"name"`
	Comment               null.String `db:"comment"`
	Alias                 null.String `db:"alias"`
	Weight                null.String `db:"weight"`
	Number                null.String `db:"number"`
	Caloric               null.Float  `db:"caloric"`
	Pfc                   null.String `db:"pfs"`
	Composition           null.String `db:"composition"`
	Image                 null.String `db:"image"`
	FullImage             null.String `db:"full_image"`
	New                   null.Bool   `db:"new"`
	Hot                   null.Bool   `db:"hot"`
	Hit                   null.Bool   `db:"hit"`
	NotDeliverySeparately bool        `db:"not_delivery_separately"`
	Volume                null.Float  `db:"volume"`
	Article               null.String `db:"article"`
	Vendor1               null.String `db:"vendor1"`
	Vendor2               null.String `db:"vendor2"`
}
type ProductItems []ProductItem

func (p *ProductItem) ToAProductItem() transport.AProductItem {
	result := transport.AProductItem{}
	result.ID = p.ID
	result.Type = p.Type
	result.Name = p.Name
	result.Comment = p.Comment.ValueOrZero()
	result.Alias = p.Alias.ValueOrZero()
	result.Weight = p.Weight.ValueOrZero()
	result.Number = p.Number.ValueOrZero()
	result.Caloric = float32(p.Caloric.ValueOrZero())
	result.Pfc = p.Pfc.ValueOrZero()
	result.Composition = p.Pfc.ValueOrZero()
	result.Image = p.Image.ValueOrZero()
	result.FullImage = p.FullImage.ValueOrZero()
	result.New = p.New.ValueOrZero()
	result.Hot = p.Hot.ValueOrZero()
	result.Hit = p.Hit.ValueOrZero()
	result.NotDeliverySeparately = p.NotDeliverySeparately
	result.Volume = float32(p.Volume.ValueOrZero())
	result.Article = p.Article.ValueOrZero()
	result.Vendor1 = base.StringInt(p.Vendor1.ValueOrZero())
	result.Vendor2 = base.StringInt(p.Vendor2.ValueOrZero())
	return result
}
