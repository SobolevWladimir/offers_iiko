package pmodifiers

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
								"id" uuid NOT NULL PRIMARY KEY, 
								"product" uuid NOT NULL REFERENCES "public"."` + tablename.OrderProduct + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,
								"modifier" uuid NOT NULL REFERENCES "public"."` + tablename.Modifications + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,
								"quantity" float NOT NULL DEFAULT 0,
                "deleted" bool DEFAULT false NOT NULL 
        );`
	tx := db.MustBegin()
	tx.MustExec(request)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."product" IS 'идентификатор продукта в заказе'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."modifier" IS 'идентификатор модификатора '`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."quantity" IS 'количество '`)
	tx.Commit()
}
