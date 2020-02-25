package products

import (
	"log"
)

func (ob *Object) CheckStructure() {
	err := UpdataStorage()
	if err != nil {
		panic(" не могу загрузить  продукты  из базы данных")
	}
}

func Fatal(funcName string, err error) {
	log.Fatal("service/webhook   in func "+funcName+" eror message", err)
}
