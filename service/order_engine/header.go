package order_engine

import (
	"altegra_offers/config"
	"altegra_offers/service/order_engine/address"
	"altegra_offers/service/order_engine/markers"
	"altegra_offers/service/order_engine/mlink"
	"altegra_offers/service/order_engine/orders"
	"altegra_offers/service/order_engine/payment"
	"altegra_offers/service/order_engine/pmodifiers"
	"altegra_offers/service/order_engine/products"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var setting config.Core
var dataBase *sqlx.DB

type Object struct {
	OrdersService     *orders.Object
	MarkerService     *markers.Object
	AddressService    *address.Object
	ProductsService   *products.Object
	PModifierService  *pmodifiers.Object
	PaymentService    *payment.Object
	MarkerLinkService *mlink.Object
}

func New() *Object {
	result := new(Object)
	result.OrdersService = orders.New()
	result.MarkerService = markers.New()
	result.AddressService = address.New()
	result.ProductsService = products.New()
	result.PModifierService = pmodifiers.New()
	result.PaymentService = payment.New()
	result.MarkerLinkService = mlink.New()
	return result
}

func (ob *Object) SetConfig(conf *config.Config) {
	ob.OrdersService.SetConfig(conf)
	ob.MarkerService.SetConfig(conf)
	ob.AddressService.SetConfig(conf)
	ob.AddressService.SetConfig(conf)
	ob.PModifierService.SetConfig(conf)
	ob.PaymentService.SetConfig(conf)
	ob.MarkerLinkService.SetConfig(conf)
	setting = conf.Core
}
func (ob *Object) SetDataBaseManager(d *sqlx.DB) {
	ob.OrdersService.SetDataBaseManager(d)
	ob.MarkerService.SetDataBaseManager(d)
	ob.AddressService.SetDataBaseManager(d)
	ob.ProductsService.SetDataBaseManager(d)
	ob.PModifierService.SetDataBaseManager(d)
	ob.PaymentService.SetDataBaseManager(d)
	ob.MarkerLinkService.SetDataBaseManager(d)
	dataBase = d
}

func connect() *sqlx.DB {
	return dataBase
}
