package orders

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
                "id" uuid NOT NULL PRIMARY KEY ,
								"point" uuid NOT NULL REFERENCES "public"."` + tablename.Point + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"client" text REFERENCES  "public"."` + tablename.Client + `" ("phone") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"parent" uuid REFERENCES "public"."` + GetTableName() + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION,
								"status" text NOT NULL REFERENCES "public"."` + tablename.Status + `" ("code") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"type" int NOT NULL DEFAULT 0, 
						    "delivery"	 int NOT NULL DEFAULT 0, 
								"person" int NOT NULL DEFAULT 1, 
								"cook_in_date" date, 
								"cook_in_time" time, 
								"paid" bool NOT NULL DEFAULT false, 
								"local_number" text, 
								"comment" text, 
								"pre_amount" float NOT NULL DEFAULT 0, 
								"amount" float NOT NULL DEFAULT 0, 
								"person_in_charge" uuid REFERENCES "public"."` + tablename.User + `" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION, 
								"offers" jsonb, 
								"offers_event" jsonb, 
								"cart" jsonb, 
								"coupon" text, 
								"last_update" timestamptz NOT NULL default current_timestamp, 
								"created" timestamptz NOT NULL default current_timestamp, 
                "deleted" bool DEFAULT false NOT NULL 
        );`
	tx := db.MustBegin()
	tx.MustExec(request)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."point" IS 'точка обработки заказа'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."parent" IS 'к какому заказу прекрепить этот заказ'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."status" IS 'статус заказа'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."type" IS '0-продажа 1- возврат '`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."delivery" IS '0-самовывоз 1-доставка 2-в зале'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."person" IS 'на какое кол-во персон готовить'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."cook_in_date" IS 'к какой дате приготовить'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."cook_in_time" IS 'к какому времени приготовить'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."paid" IS ' оплачен ли товар'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."local_number" IS 'локальный номер заказа'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."comment" IS 'Комментарий'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."pre_amount" IS 'сумма заказа без учета  скидок'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."amount" IS 'итоговая сумма заказ'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."person_in_charge" IS 'отвестсвенное лицо'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."offers" IS 'сработанные акции'`)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."last_update" IS 'последнее обновление данных'`)
	tx.Commit()
}
