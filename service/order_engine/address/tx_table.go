package address

import "github.com/jmoiron/sqlx"

func TxInsert(tx *sqlx.Tx, entity *Address) error {
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					"order", 
					city, 
					street_id, 
					street_text,
					building, 
					entrance, 
					floor, 
					room, 
					comment
	) VALUES (
					:order, 
					:city, 
					:street_id, 
					:street_text,
					:building, 
					:entrance, 
					:floor, 
					:room, 
					:comment
					
	)`
	_, err := tx.NamedExec(reqInsert, entity)
	return err
}
func TxSave(tx *sqlx.Tx, entity *Address) error {
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					city=:city, 
					street_id=:street_id, 
					street_text=:street_text,
					building=:building, 
					entrance=:entrance, 
					floor=:floor, 
					room=:room, 
					comment=:comment,
					deleted='false'
					WHERE "order"=:order`
	_, err := tx.NamedExec(reqUpdate, entity)
	return err
}
func TxSaveSafety(tx *sqlx.Tx, entity *Address) error {
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					"order", 
					city, 
					street_id, 
					street_text,
					building, 
					entrance, 
					floor, 
					room, 
					comment
	) VALUES (
					:order, 
					:city, 
					:street_id, 
					:street_text,
					:building, 
					:entrance, 
					:floor, 
					:room, 
					:comment
					) ON CONFLICT ("order") DO  UPDATE SET  
					city=:city, 
					street_id=:street_id, 
					street_text=:street_text,
					building=:building, 
					entrance=:entrance, 
					floor=:floor, 
					room=:room, 
					comment=:comment,
					deleted='false'`
	_, err := tx.NamedExec(reqInsert, entity)
	return err
}
func TxRemoveByOrderId(tx *sqlx.Tx, order string) error {
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					deleted='true'
					WHERE "order"=$1`
	_, err := tx.Queryx(reqUpdate, order)
	return err

}
