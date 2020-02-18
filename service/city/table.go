package city

func FindAll() (Citys, error) {
	result := Citys{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where deleted='false'`)
	return result, err

}
func ExistById(id string) (bool, error) {
	var result bool
	db := connect()
	err := db.Get(&result, `SELECT EXISTS(SELECT * FROM "public"."`+GetTableName()+`" where id='`+id+`')`)
	if err != nil {
		return result, err
	}
	return result, nil
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
func FindOneById(id string) (City, error) {
	result := City{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where id='`+id+`'`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneByName(name string) (City, error) {
	result := City{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where name='`+name+`' and deleted='false'`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *City) error {
	db := connect()

	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (id, name) VALUES (:id, :name)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *City) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET name=:name, deleted='false' WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveById(id string) {
	db := connect()
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id='` + id + `'`
	db.MustExec(request)
}
