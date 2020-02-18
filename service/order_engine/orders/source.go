package orders

import (
	"altegra_offers/lib/tablename"
	"log"
)

func (ob *Object) CheckStructure() {
	db := connect()
	//defer db.Close()
	if !checkDb(db) {
		createTable(db)
	}
}

func Fatal(funcName string, err error) {
	log.Fatal("service/order_engine/orders   in func "+funcName+" eror message", err)
}
func GetTableName() string {
	return tablename.Orders
}
