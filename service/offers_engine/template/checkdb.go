package template

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createTable(db *sqlx.DB) {
	var request = `CREATE TABLE IF NOT EXISTS ` + GetTableName() + ` (
                id int(11) NOT NULL PRIMARY KEY auto_increment,
								setrule json NOT NULL, 
                sort int(11) DEFAULT 0 NOT NULL						
        );`
	db.MustExec(request)
}
