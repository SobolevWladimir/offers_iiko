package webhook

import (
	"log"
	"offers_iiko/lib/tablename"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func checkDb(db *sqlx.DB) bool {
	var request = `select * from pg_tables where tablename='` + tablename.Webhook + `';`
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
	var request = `CREATE TABLE "public"."` + tablename.Webhook + `" (
						"id" uuid NOT NULL,
						"type" uuid NOT NULL,
						"url" text NOT NULL,
            "deleted" bool DEFAULT false NOT NULL
						);
						 ALTER TABLE "public"."` + tablename.Webhook + `" ADD PRIMARY KEY ("id");`

	db.MustExec(request)

}
