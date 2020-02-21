package products

import (
	"log"
)

func (ob *Object) CheckStructure() {
}

func Fatal(funcName string, err error) {
	log.Fatal("service/webhook   in func "+funcName+" eror message", err)
}
