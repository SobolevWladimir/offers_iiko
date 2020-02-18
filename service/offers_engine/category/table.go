package category

func FindAll() (Categorys, error) {
	result := Categorys{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM `+GetTableName()+` `)
	return result, err

}
func FindByCity(city int) (Categorys, error) {
	result := Categorys{}
	db := connect()

	err := db.Select(&result, `SELECT * FROM `+GetTableName()+` where city=?`, city)
	return result, err

}
func FindOneById(id string) (Category, error) {
	result := Category{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM `+GetTableName()+` where id=?`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneByName(name string) (Category, error) {
	result := Category{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM `+GetTableName()+` where name=? `, name)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Category) error {
	db := connect()

	var reqInsert = `INSERT INTO ` + GetTableName() + ` (name, parent, city,  sort) VALUES ( :name, :parent, :city,  :sort)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Category) error {
	db := connect()
	var reqUpdate = `UPDATE ` + GetTableName() + ` SET name=:name, parent=:parent, city=:city,  sort=:sort  WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveById(id int) {
	db := connect()
	var request = `DELETE FROM  ` + GetTableName() + ` WHERE id=?`
	db.MustExec(request, id)
}
