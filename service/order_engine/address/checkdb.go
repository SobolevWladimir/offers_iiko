package address

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
								"order" uuid NOT NULL UNIQUE REFERENCES "public"."` + tablename.Orders + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"city" uuid NOT NULL REFERENCES "public"."` + tablename.City + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"street_id" uuid  REFERENCES  "public"."` + tablename.Street + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"street_text" text, 
								"building" text, 
								"entrance" text, 
								"floor" text, 
								"room" text, 
								"comment" text, 
                "deleted" bool DEFAULT false NOT NULL 
        );`
	tx := db.MustBegin()
	tx.MustExec(request)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."order" IS 'идентификатор заказа'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."city" IS 'идентификатор города'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."street_id" IS 'идентификатор улицы'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."street_text" IS 'если улицы нет в базе данных то записываем улицу текстом'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."building" IS 'здание (дом)'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."entrance" IS 'подъезд'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."floor" IS 'этаж'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."room" IS 'комната(квартира)'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."comment" IS 'комментарий'`)
	tx.Commit()
}
