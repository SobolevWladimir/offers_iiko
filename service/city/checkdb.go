package city

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const DefaultId = "1e7e9538-5475-4cc8-8567-6b5907494956"
const DefaulName = "Смоленск"

func createTable(db *sqlx.DB) {
	var request = `CREATE TABLE IF NOT EXISTS "` + GetTableName() + `" (
						"id" uuid NOT NULL PRIMARY KEY,
						"name" text NOT NULL, 						
						"deleted" bool DEFAULT false NOT NULL
						);`
	db.MustExec(request)
	tx := db.MustBegin()
	tx.Commit()
}
