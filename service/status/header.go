package status

import (
	"altegra_offers/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var setting config.Core
var dataBase *sqlx.DB

type Object struct {
}

func New() *Object {
	return new(Object)
}

func (ob *Object) SetConfig(conf *config.Config) {
	setting = conf.Core
}
func (ob *Object) SetDataBaseManager(d *sqlx.DB) {
	dataBase = d
}

func connect() *sqlx.DB {
	return dataBase
}
