package coupon

import (
	"altegra_offers/lib/tablename"
	"log"
)

func (ob *Object) CheckStructure() {
	db := connect()
	createTable(db)
}

func Fatal(funcName string, err error) {
	log.Fatal("service/coupon   in func "+funcName+" eror message", err)
}
func GetTableName() string {
	return tablename.Coupon
}
