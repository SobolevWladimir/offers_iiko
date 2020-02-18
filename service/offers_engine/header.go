package offers_engine

import (
	"altegra_offers/config"
	"altegra_offers/service/offers_engine/category"
	"altegra_offers/service/offers_engine/offer"
	"altegra_offers/service/offers_engine/template"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var setting config.Config
var dataBase *sqlx.DB

type Object struct {
	CategoryService *category.Object
	OffersService   *offer.Object
	TemplateService *template.Object
}

func New() *Object {
	result := new(Object)
	result.CategoryService = category.New()
	result.OffersService = offer.New()
	result.TemplateService = template.New()
	return result
}

func (ob *Object) SetConfig(conf *config.Config) {
	ob.CategoryService.SetConfig(conf)
	ob.OffersService.SetConfig(conf)
	ob.TemplateService.SetConfig(conf)
	setting = *conf
}
func (ob *Object) SetDataBaseManager(d *sqlx.DB) {
	ob.CategoryService.SetDataBaseManager(d)
	ob.OffersService.SetDataBaseManager(d)
	ob.TemplateService.SetDataBaseManager(d)
	dataBase = d
}

func connect() *sqlx.DB {
	return dataBase
}
