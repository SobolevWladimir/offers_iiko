package street

import (
	"altegra_offers/lib/tablename"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DefaultType = "Улица"
const DefaultTypeShort = "ул"

func checkDb(db *sqlx.DB) bool {
	var request = `select * from pg_tables where tablename='` + GetTableName() + `';`
	rows, err := db.Queryx(request)
	if err != nil {
		log.Fatal(request, err)
	}
	defer rows.Close()
	result := rows.Next()

	if !result {
		return false
	}
	return true

}
func createTable(db *sqlx.DB) {
	var request = `CREATE TABLE "public"."` + GetTableName() + `" (
						"id" uuid NOT NULL PRIMARY KEY , 
						"name" text NOT NULL, 						
						"type" text DEFAULT '` + DefaultType + `' NOT NULL, 						
						"type_short" text DEFAULT '` + DefaultTypeShort + `'  NOT NULL, 						
						"city" uuid NOT NULL REFERENCES "public"."` + tablename.City + `" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION,
						"deleted" bool DEFAULT false NOT NULL, 
								CONSTRAINT pk_city_street UNIQUE (name,city)
						);
						 `
	db.MustExec(request)

}
