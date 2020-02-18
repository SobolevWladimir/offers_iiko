package offer

func FindAll() (Offers, error) {
	result := Offers{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM `+GetTableName()+` `)
	return result, err

}
func FindByCategory(category int) (Offers, error) {
	result := Offers{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM `+GetTableName()+` where category=? `, category)
	return result, err

}
func FindOneById(id string) (Offer, error) {
	result := Offer{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM `+GetTableName()+` where id=$1`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneByName(name string) (Offer, error) {
	result := Offer{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM `+GetTableName()+` where name=$1 and deleted='false'`, name)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Offer) error {
	db := connect()

	var reqInsert = `INSERT INTO ` + GetTableName() + ` (
					id,
					active, 
					name, 
					status, 
					algorithm, 
					setrules, 
					actions, 
					category, 
					sort
	) VALUES (
					:id,
					:active, 
					:name, 
					:status, 
					:algorithm, 
					:setrules, 
					:actions, 
					:category, 
					:sort
	)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Offer) error {
	db := connect()
	var reqUpdate = `UPDATE ` + GetTableName() + ` SET 
					active=:active, 
					name=:name, 
					status=:status, 
					algorithm=:algorithm, 
					setrules=:setrules, 
					actions=:actions, 
					category=:category, 
					sort=:sort
	WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveById(id int) {
	db := connect()
	var request = `DELETE FROM  ` + GetTableName() + ` WHERE id=?`
	db.MustExec(request, id)
}
