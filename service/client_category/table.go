package client_category

func FindAll() (ClientCategorys, error) {
	result := ClientCategorys{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`"`)
	return result, err
}
func FindOneById(id string) (ClientCategory, error) {
	result := ClientCategory{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where id=$1`, id)
	return result, err
}

func FindOneByName(name string) (ClientCategory, error) {
	result := ClientCategory{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where name=$1 and deleted='false'`, name)
	return result, err
}

func Save(client *ClientCategory) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() +
		`" SET
				name=:name,
				deleted='false' WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, &client)
	return err
}

func RemoveById(id string) {
	db := connect()
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id='` + id + `'`
	db.MustExec(request)

}
func Insert(client *ClientCategory) error {
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" 
		(id, name)	VALUES
		 (:id, :name )`
	_, err := db.NamedExec(reqInsert, &client)
	return err
}
