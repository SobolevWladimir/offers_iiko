package offer

import (
	"altegra_offers/lib/tablename"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func createTable(db *sqlx.DB) {
	var request = `CREATE TABLE  IF NOT EXISTS ` + GetTableName() + ` (
                id int(11) NOT NULL PRIMARY KEY auto_increment,
								active bool DEFAULT false NOT NULL, 
                name text NOT NULL, 						
								status text  NOT NULL, 
								algorithm int DEFAULT 0 NOT NULL, 
								setrules json, 
								actions json, 
                category int NOT NULL ,
                sort int DEFAULT 0 NOT NULL,						
								FOREIGN KEY (category) REFERENCES ` + tablename.OffersCategory + ` (id) ON DELETE NO ACTION ON UPDATE NO ACTION 
        );`
	db.MustExec(request)
}
