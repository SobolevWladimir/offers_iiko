package client

import (
	"altegra_offers/lib/base"
	"altegra_offers/lib/tablename"

	"github.com/huandu/go-sqlbuilder"
)

func FindAll() (Clients, error) {
	result := Clients{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+tablename.Client+`" where deleted='false'`)
	return result, err
}
func FindAllByPages(page, limit int) (Clients, error) {
	result := Clients{}
	db := connect()
	var offset = 0
	if page > 0 {
		offset = (page - 1) * limit
	}
	err := db.Select(&result, `SELECT * FROM "public"."`+tablename.Client+`" where  deleted='false' LIMIT $1 OFFSET $2`, limit, offset)
	return result, err
}
func FindOneByName(name string) (Client, error) {
	result := Client{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+tablename.Client+`" where name=$1 and deleted='false'`, name)

	return result, err

}
func FindOneByPhone(phone string) (Client, error) {
	result := Client{}
	result.Bonuses = 0
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+tablename.Client+`" where phone=$1  and deleted='false'`, phone)
	return result, err

}
func FindByPhone(page int, limit int, phone string) (Clients, error) {
	result := Clients{}
	var search = "%" + phone + "%"
	var offset = 0
	if page > 0 {
		offset = (page - 1) * limit
	}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+tablename.Client+`" where phone LIKE $1 and deleted='false' LIMIT $2 OFFSET $3`, search, limit, offset)
	return result, err

}
func FindByPhones(phones []string) (Clients, error) {
	result := Clients{}

	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("*")
	sb.From(tablename.Client)
	sb.Where(sb.In("phone", base.StringToInterface(phones)...))
	sql, args := sb.Build()
	db := connect()
	err := db.Select(&result, sql, args...)
	return result, err
}
func FindByCard(card string) ([]Client, error) {
	result := []Client{}
	db := connect()

	err := db.Select(&result, `SELECT * FROM "public"."`+tablename.Client+`" where card_number LIKE '%`+card+`%' and deleted='false' LIMIT 10`)

	return result, err

}

func Save(client *Client) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + tablename.Client +
		`" SET 
		 	name=:name,
		   	last_name=:lastname,
		    email=:email,
		    sex=:sex,
		    birth_date=:birth_date,
		    bonuses=:bonuses,
			deleted='false' WHERE phone=:phone`
	_, err := db.NamedExec(reqUpdate, &client)
	return err
}
func AppendBonuses(phone string, bonuse float32) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + tablename.Client +
		`" SET 
		    bonuses=bonuses+ $2,
			deleted='false' WHERE phone=$1`
	_, err := db.Query(reqUpdate, phone, bonuse)
	return err
}

func RemoveByPhone(phone string) error {
	db := connect()
	var sql = `Update "public"."` + tablename.Client + `" SET deleted='true' WHERE phone=$1`
	_, err := db.Query(sql, phone)
	return err
}
func Insert(client *Client) error {
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + tablename.Client + `" 
		(phone, name, last_name, email, sex, birth_date, bonuses)	VALUES
		 (:phone, :name, :lastname, :email, :sex, :birth_date, :bonuses)`
	_, err := db.NamedExec(reqInsert, client)
	return err
}
