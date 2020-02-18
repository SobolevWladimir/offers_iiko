package order_engine

import (
	"altegra_offers/service/order_engine/address"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
)

type Address struct {
	City       string      `json:"city"`
	StreetId   null.String `json:"street_id"`
	StreetText null.String `json:"street_text"`
	Building   null.String `json:"building"`
	Entrance   null.String `json:"entrance"`
	Floor      null.String `json:"floor"`
	Room       null.String `json:"room"`
	Comment    null.String `json:"comment"`
}

func toAddress(ad *address.Address) Address {
	result := Address{}
	result.City = ad.City
	result.StreetId = ad.StreetId
	result.StreetText = ad.StreetText
	result.Building = ad.Building
	result.Entrance = ad.Entrance
	result.Floor = ad.Floor
	result.Room = ad.Room
	result.Comment = ad.Comment
	return result
}
func (ad *Address) toDBEntity(order string) address.Address {
	result := address.Address{}
	result.Order = order
	result.City = ad.City
	result.StreetId = ad.StreetId
	result.StreetText = ad.StreetText
	result.Building = ad.Building
	result.Entrance = ad.Entrance
	result.Floor = ad.Floor
	result.Room = ad.Room
	result.Comment = ad.Comment
	return result

}
func FindAddressByOrder(order string) (Address, error) {
	ad, err := address.FindOneByOrder(order)
	return toAddress(&ad), err
}
func (ad *Address) Insert(order string) error {
	entity := ad.toDBEntity(order)
	return address.Insert(&entity)
}
func (ad *Address) Save(order string) error {
	entity := ad.toDBEntity(order)
	return address.SaveSafety(&entity)
}
func (ad *Address) TxSave(tx *sqlx.Tx, order string) error {
	entity := ad.toDBEntity(order)
	return address.TxSaveSafety(tx, &entity)
}
