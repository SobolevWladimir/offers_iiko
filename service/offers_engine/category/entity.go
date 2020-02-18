package category

import (
	"gopkg.in/guregu/null.v3"
)

type Category struct {
	Id     int      `db:"id" json:"id"`
	Name   string   `db:"name" json:"name"`
	Parent null.Int `db:"parent" json:"parent"`
	Sort   int      `db:"sort" json:"sort"`
	City   int      `db:"city" json:"city"`
}
type Categorys []Category
