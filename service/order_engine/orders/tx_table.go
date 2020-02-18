package orders

import "github.com/jmoiron/sqlx"

func TxInsert(tx *sqlx.Tx, entity *Order) error {
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
	_, err := tx.NamedExec(reqInsert, entity)
	return err
}
func TxSave(tx *sqlx.Tx, entity *Order) error {
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
	_, err := tx.NamedExec(reqUpdate, entity)
	return err
}
