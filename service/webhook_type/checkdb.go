package webhook_type

import (
	"altegra_offers/lib/tablename"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func checkDb(db *sqlx.DB) bool {
	var request = `select * from pg_tables where tablename='` + tablename.Webhook_type + `';`
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
	var request = `CREATE TABLE "public"."` + tablename.Webhook_type + `" (
						"id" uuid NOT NULL,
						"alias" text NOT NULL,
						"name" text NOT NULL,
						"comment" text,
            "deleted" bool DEFAULT false NOT NULL
						);
						 ALTER TABLE "public"."` + tablename.Webhook_type + `" ADD PRIMARY KEY ("id");`

	db.MustExec(request)

}
