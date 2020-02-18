package orders

func FindAll(date_start, date_end string) (Orders, error) {
	result := Orders{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where (last_update BETWEEN $1 and $2) and deleted='false' ORDER BY created;`, date_start, date_end)
	return result, err

}
func FindAllByPoint(date_start, date_end, point string) (Orders, error) {
	result := Orders{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where (last_update BETWEEN $1 and $2) and point=$3 and deleted='false' ORDER BY created;`, date_start, date_end, point)
	return result, err

}

func FindOneById(id string) (Order, error) {
	result := Order{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where id=$1`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Order) error {
	db := connect()

	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
	id,
	point,
	client,
	parent,
	status,
	type,
	delivery,
	person,
	cook_in_date,
	cook_in_time,
	paid,
	local_number,
	comment,
	pre_amount,
	amount,
	person_in_charge,
	offers, 
	offers_event,
  cart
	) VALUES (
	:id,
	:point,
	:client,
	:parent,
	:status,
	:type,
	:delivery,
	:person,
	:cook_in_date,
	:cook_in_time,
	:paid,
	:local_number,
	:comment,
	:pre_amount,
	:amount,
	:person_in_charge,
	:offers, 
	:offers_event, 
	:cart
	)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Order) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
	point=:point,
	client=:client,
	parent=:parent,
	status=:status,
	type=:type,
	delivery=:delivery,
	person=:person,
	cook_in_date=:cook_in_date,
	cook_in_time=:cook_in_time,
	paid=:paid,
	local_number=:local_number,
	comment=:comment,
	pre_amount=:pre_amount,
	amount=:amount,
	person_in_charge=:person_in_charge,
	offers=:offers,
	offers_event=:offers_event,
	cart=:cart, 
	deleted='false' WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func SaveOfferEvents(entity *Order) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
	offers_event=:offers_event,
	deleted='false' WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveById(id string) error {
	db := connect()
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id=$1`
	_, err := db.Queryx(request, id)
	return err
}
