package mlink

import (
	"altegra_offers/lib/tablename"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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
								"order" uuid NOT NULL REFERENCES "public"."` + tablename.Orders + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"marker" text NOT NULL REFERENCES "public"."` + tablename.OrderMarker + `" ("name") ON DELETE NO ACTION ON UPDATE NO ACTION,
                "deleted" bool DEFAULT false NOT NULL,
								CONSTRAINT un_OrderMarkerLinks UNIQUE ("order",marker)
        );`
	tx := db.MustBegin()
	tx.MustExec(request)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."order" IS 'идентификатор заказа'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."marker" IS 'Маркер заказа'`)
	tx.Commit()
}
