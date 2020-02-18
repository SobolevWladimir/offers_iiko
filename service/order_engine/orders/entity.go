package orders

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

type Order struct {
	Id             string      `db:"id"`
	Point          string      `db:"point"`
	Client         null.String `db:"client"`
	Parent         null.String `db:"parent"`
	Status         string      `db:"status"`
	Type           int         `db:"type"`
	Delivery       int         `db:"delivery"`
	Person         int         `db:"person"`
	CookInDate     time.Time   `db:"cook_in_date"`
	CookInTime     null.Time   `db:"cook_in_time"`
	Paid           bool        `db:"paid"`
	LocalNumber    null.String `db:"local_number"`
	Comment        null.String `db:"comment"`
	PreAmount      float32     `db:"pre_amount"`
	Amount         float32     `db:"amount"`
	PersonInCharge null.String `db:"person_in_charge"`
	Offers         null.String `db:"offers"`
	OfferEvent     null.String `db:"offers_event"`
	Cart           null.String `db:"cart"`
	Coupon         null.String `db:"coupon"`
	LastUpdate     time.Time   `db:"last_update"`
	Created        time.Time   `db:"created"`
}

type Orders []Order
