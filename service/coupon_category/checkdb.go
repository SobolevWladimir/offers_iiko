package coupon_category

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createTable(db *sqlx.DB) {
	var request = `CREATE TABLE IF NOT EXISTS ` + GetTableName() + ` (
                id int(11) NOT NULL PRIMARY KEY auto_increment,
                name text NOT NULL, 						
                parent int(11),  
								city int(11) NOT NULL , 
                sort int(11) DEFAULT 0 NOT NULL, 
								FOREIGN KEY (parent)REFERENCES ` + GetTableName() + ` (id) ON DELETE NO ACTION ON UPDATE NO ACTION
        );`
	db.MustExec(request)
}
