package street

import (
	"altegra_offers/lib/tablename"
	"log"
)

func (ob *Object) CheckStructure() {
	db := connect()
	if !checkDb(db) {
		createTable(db)
	}

}

func Fatal(funcName string, err error) {
	log.Fatal("service/street   in func "+funcName+" eror message", err)
}
func GetTableName() string {
	return tablename.Street
}
