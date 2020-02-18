package client_category

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
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
            "id" uuid NOT NULL UNIQUE,
						"name" text NOT NULL UNIQUE, 
             "deleted" bool DEFAULT false NOT NULL
						);`
	db.MustExec(request)
}
