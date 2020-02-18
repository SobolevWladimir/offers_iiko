package webhook_type

import "altegra_offers/lib/tablename"

func FindAll() ([]Webhook_type, error) {
	result := []Webhook_type{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+tablename.Webhook_type+`" where deleted='false'`)

	return result, err

}
