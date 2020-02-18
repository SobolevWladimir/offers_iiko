package city

import (
	"altegra_offers/lib/tablename"
	"log"
)

func (ob *Object) CheckStructure() {
	db := connect()
	//defer db.Close()
	createTable(db)

}

func Fatal(funcName string, err error) {
	log.Fatal("service/city   in func "+funcName+" eror message", err)
}
func GetTableName() string {
	return tablename.City
}
