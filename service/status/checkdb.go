package status

import (
	"altegra_offers/lib/tablename"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func checkDb(db *sqlx.DB) bool {
	var request = `select * from pg_tables where tablename='` + tablename.Status + `';`
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
	var request = `CREATE TABLE "public"."` + tablename.Status + `" (
						"name" text NOT NULL,
						"code" text NOT NULL UNIQUE,
						"priority" int DEFAULT 0 NOT NULL,
						"color" text NOT NULL,
						"comment" text, 
			"deleted" bool DEFAULT false NOT NULL
						);
						 `
	db.MustExec(request)
	if err := Insert(&StatusNewOrder); err != nil {
		log.Fatal("not insert status", err)
	}
	Insert(&StatusAwaitOrder)
	Insert(&StatusPrepareOrder)
	Insert(&StatusCookedOrder)
	Insert(&StatusWaitingSendOrder)
	Insert(&StatusShippedOrder)
	Insert(&StatusDeliveredOrder)
	Insert(&StatusDeliveredOrder)
	Insert(&StatusClosedOrder)
	Insert(&StatusCancelOrder)
}
