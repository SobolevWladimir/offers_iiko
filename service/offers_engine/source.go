package offers_engine

import (
	"log"
)

func (ob *Object) CheckStructure() {
	ob.CategoryService.CheckStructure()
	ob.OffersService.CheckStructure()
	ob.TemplateService.CheckStructure()
}

func Fatal(funcName string, err error) {
	log.Fatal("service/offers_engine   in func "+funcName+" eror message", err)
}
