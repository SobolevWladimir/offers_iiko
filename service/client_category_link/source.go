package client_category_link

import (
	"altegra_offers/lib/tablename"
	"log"
)

func (ob *Object) CheckStructure() {
	db := connect()
	//defer db.Close()

	if !checkDbUser(db) {
		createTableUser(db)
	}
	// initAccessData()
}
func Fatal(funcName string, err error) {
	log.Fatal("service/user   in func "+funcName+" error message: ", err)

}
func GetTableName() string {
	return tablename.ClientCategoryLink
}
