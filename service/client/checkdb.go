package client

import (
	"altegra_offers/lib/tablename"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func checkDbUser(db *sqlx.DB) bool {
	var request = `select * from pg_tables where tablename='` + tablename.Client + `';`

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
func createTableUser(db *sqlx.DB) {

	var request = `CREATE TABLE "public"."` + tablename.Client + `" (
						"phone" text NOT NULL UNIQUE, 
						"name" text NOT NULL, 
						"last_name" text,
						"email" text, 
						"sex" int DEFAULT 0 NOT NULL, 
						"birth_date" date,
						 "bonuses" float4 DEFAULT 0 NOT NULL, 
             "deleted" bool DEFAULT false NOT NULL
						);`
	db.MustExec(request)
}
