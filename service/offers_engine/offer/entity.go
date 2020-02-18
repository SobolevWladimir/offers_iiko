package offer

import (
	"gopkg.in/guregu/null.v3"
)

type Offer struct {
	Id        int         `db:"id"`
	Active    bool        `db:"active"`
	Name      string      `db:"name"`
	Status    string      `db:"status"`
	Algorithm int         `db:"algorithm"`
	SetRules  null.String `db:"setrules"`
	Actions   null.String `db:"actions"`
	Category  int         `db:"category"`
	Sort      int         `db:"sort"`
}
type Offers []Offer
