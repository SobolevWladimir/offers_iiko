package client_category_link

import (
	"altegra_offers/lib/tablename"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func checkDbUser(db *sqlx.DB) bool {
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
func createTableUser(db *sqlx.DB) {

	var request = `CREATE TABLE "public"."` + GetTableName() + `" (
						"client" text NOT NULL REFERENCES "public"."` + tablename.Client + `" ("phone") ON DELETE CASCADE ON UPDATE CASCADE,
						"category" text NOT NULL REFERENCES "public"."` + tablename.ClientCategory + `" ("name") ON DELETE CASCADE ON UPDATE CASCADE,  
             "deleted" bool DEFAULT false NOT NULL
						);`
	db.MustExec(request)
}
