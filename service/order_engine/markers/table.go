package markers

func FindAll() ([]string, error) {
	result := []string{}
	db := connect()
	err := db.Select(&result, `SELECT name FROM "public"."`+GetTableName()+`" where deleted='false'`)
	return result, err
}
func Insert(name string) error {
	db := connect()
	_, err := db.Queryx(`INSERT INTO "public"."`+GetTableName()+`" (name) VALUES ($1)`, name)
	return err
}
func RemoveByName(name string) error {
	db := connect()
	_, err := db.Queryx(`Update "public"."`+GetTableName()+`" SET deleted='true' WHERE name=$1`, name)
	return err
}
func ExistByName(name string) (bool, error) {
	var result bool
	db := connect()
	err := db.Get(&result, `SELECT EXISTS(SELECT * FROM "public"."`+GetTableName()+`" where name='`+name+`')`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Save(name string) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET deleted='false' WHERE name=$1`
	_, err := db.Queryx(reqUpdate, name)
	return err
}
