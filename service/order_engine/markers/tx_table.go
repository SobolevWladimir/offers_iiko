package markers

import "github.com/jmoiron/sqlx"

func TxRemoveByName(tx *sqlx.Tx, name string) error {
	_, err := tx.Queryx(`Update "public"."`+GetTableName()+`" SET deleted='true' WHERE name=$1`, name)
	return err
}
func TxExistByName(tx *sqlx.Tx, name string) (bool, error) {
	var result bool
	err := tx.Get(&result, `SELECT EXISTS(SELECT * FROM "public"."`+GetTableName()+`" where name='`+name+`')`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func TxSave(tx *sqlx.Tx, name string) error {
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET deleted='false' WHERE name=$1`
	_, err := tx.Queryx(reqUpdate, name)
	return err
}
