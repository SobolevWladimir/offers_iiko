package address

func FindAll() (Addresses, error) {
	result := Addresses{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where deleted='false'`)
	return result, err
}
func FindAllByOrder(order string) (Addresses, error) {
	result := Addresses{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where "order"=$1 and deleted='false'`, order)
	return result, err
}
func FindOneById(id string) (Address, error) {
	result := Address{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where id=$1`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneByOrder(id string) (Address, error) {
	result := Address{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where "order"=$1 and deleted='false'`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Address) error {
	db := connect()
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
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Address) error {
	db := connect()
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
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func SaveSafety(entity *Address) error {
	db := connect()
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
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func RemoveByOrderId(order string) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					deleted='true'
					WHERE "order"=$1`
	_, err := db.Queryx(reqUpdate, order)
	return err

}
