package offer

import (
	"altegra_offers/lib/tablename"
	"log"
)

func (ob *Object) CheckStructure() {
	db := connect()
	createTable(db)
}

func Fatal(funcName string, err error) {
	log.Fatal("service/coupon_offers   in func "+funcName+" eror message", err)
}
func GetTableName() string {
	return tablename.Offers
}
