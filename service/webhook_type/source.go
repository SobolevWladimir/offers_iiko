package webhook_type

import (
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
	log.Fatal("service/point   in func "+funcName+" eror message", err)
}