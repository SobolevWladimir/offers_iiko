package setting

import (
	"log"
	"offers_iiko/lib/tablename"
)

func (ob *Object) CheckStructure() {
	err := UpdataStorage()
	if err != nil {
		panic(" не могу загрузить настройки iiko  из базы данных")
	}
}

func Fatal(funcName string, err error) {
	log.Fatal("service/setting   in func "+funcName+" eror message", err)
}
func GetTableName() string {
	return tablename.Iiko
}
