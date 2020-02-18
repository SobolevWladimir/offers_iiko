package service

import (
	"altegra_offers/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var setting *config.Config
var services []ServiceInterface

type ServiceInterface interface {
	SetConfig(*config.Config)
	SetDataBaseManager(d *sqlx.DB)
	CheckStructure()
}

func SetConfig(conf *config.Config) {
	setting = conf
	db := connect(conf.Core.DbDriver, conf.Core.DbSourceName)
	//@todo  вынести в конфиг
	db.SetMaxOpenConns(100)
	for _, sv := range services {
		sv.SetConfig(conf)
		sv.SetDataBaseManager(db)
	}
}

func CheckStructure() {
	for _, sv := range services {
		sv.CheckStructure()
	}
}
func addService(ser ServiceInterface) {
	services = append(services, ser)
}
func connect(driver string, source_name string) *sqlx.DB {
	db, err := sqlx.Connect(driver, source_name)
	if err != nil {
		log.Fatal("service/service.go func connect" + err.Error())
	}
	return db.Unsafe()
}
