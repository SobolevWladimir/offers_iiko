package webhook

import (
	"log"
)

func (ob *Object) CheckStructure() {
	//db := connect()
	// if !checkDb(db) {
	//   createTable(db)
	// }
	// UpdateBuffer()
}

func Fatal(funcName string, err error) {
	log.Fatal("service/webhook   in func "+funcName+" eror message", err)
}
