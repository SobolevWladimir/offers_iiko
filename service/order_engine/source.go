package order_engine

import (
	"log"
)

func (ob *Object) CheckStructure() {
	ob.OrdersService.CheckStructure()
	ob.MarkerService.CheckStructure()
	ob.AddressService.CheckStructure()
	ob.ProductsService.CheckStructure()
	ob.PModifierService.CheckStructure()
	ob.PaymentService.CheckStructure()
	ob.MarkerLinkService.CheckStructure()
}

func Fatal(funcName string, err error) {
	log.Fatal("service/offers_engine   in func "+funcName+" eror message", err)
}
