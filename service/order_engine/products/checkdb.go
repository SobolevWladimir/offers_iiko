package products

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
								"order" uuid NOT NULL REFERENCES "public"."` + tablename.Orders + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"product" uuid NOT NULL REFERENCES "public"."` + tablename.Product + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,
								"quantity" float NOT NULL DEFAULT 0,
								"comment" text, 
                "deleted" bool DEFAULT false NOT NULL 
        );`
	tx := db.MustBegin()
	tx.MustExec(request)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."order" IS 'идентификатор заказа'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."product" IS 'идентификатор продукта'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."quantity" IS 'количество '`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."comment" IS 'комментарий '`)
	tx.Commit()
}
