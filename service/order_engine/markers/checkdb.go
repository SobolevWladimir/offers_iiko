package markers

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
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
								"name" text NOT NULL UNIQUE, 
                "deleted" bool DEFAULT false NOT NULL 
        );`
	tx := db.MustBegin()
	tx.MustExec(request)
	tx.MustExec(`COMMENT ON COLUMN "public"."` + GetTableName() + `"."name" IS 'Название маркера'`)
	tx.MustExec(`INSERT INTO "public"."` + GetTableName() + `" (name) VALUES ('android')`)
	tx.MustExec(`INSERT INTO "public"."` + GetTableName() + `" (name) VALUES ('ios')`)
	tx.MustExec(`INSERT INTO "public"."` + GetTableName() + `" (name) VALUES ('сайт')`)
	tx.MustExec(`INSERT INTO "public"."` + GetTableName() + `" (name) VALUES ('мобильное приложение')`)
	tx.MustExec(`INSERT INTO "public"."` + GetTableName() + `" (name) VALUES ('через точку')`)
	tx.Commit()
}
