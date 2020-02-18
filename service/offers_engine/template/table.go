package template

func FindAll() (Templates, error) {
	result := Templates{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where deleted='false' ORDER BY sort`)
	return result, err

}
func FindOneById(id string) (Template, error) {
	result := Template{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where id=$1`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneByName(name string) (Template, error) {
	result := Template{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where name=$1 and deleted='false'`, name)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Template) error {
	db := connect()

	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id, 
					setrule,
					sort
	) VALUES (
					:id,
					:setrule, 
					:sort 
	)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Template) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					setrule=:setrule, 
					sort=:sort, 
	        deleted='false'
	WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveById(id string) error {
	db := connect()
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id=$1`
	_, err := db.Queryx(request, id)
	return err
}
