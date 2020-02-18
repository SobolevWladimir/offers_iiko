package coupon

import (
	"altegra_offers/lib/tablename"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createTable(db *sqlx.DB) {
	var request = `CREATE TABLE  IF NOT EXISTS ` + GetTableName() + ` (
                id int(11) NOT NULL PRIMARY KEY auto_increment,
                name text NOT NULL, 						
								status bool DEFAULT false NOT NULL, 
								type int(11) DEFAULT 0 NOT NULL, 
								comment longtext, 
                sort int(11) DEFAULT 0 NOT NULL,						
								category int(11) NOT NULL, 
								FOREIGN KEY (category) REFERENCES ` + tablename.CouponCategory + ` (id) ON DELETE NO ACTION ON UPDATE NO ACTION 
        );`
	db.MustExec(request)
}
