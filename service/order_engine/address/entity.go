package address

import (
	"gopkg.in/guregu/null.v3"
)

type Address struct {
	Order      string      `db:"order"`
	City       string      `db:"city"`
	StreetId   null.String `db:"street_id"`
	StreetText null.String `db:"street_text"`
	Building   null.String `db:"building"`
	Entrance   null.String `db:"entrance"`
	Floor      null.String `db:"floor"`
	Room       null.String `db:"room"`
	Comment    null.String `db:"comment"`
}
type Addresses []Address
