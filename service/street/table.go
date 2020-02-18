package street

import (
	"altegra_offers/lib/tablename"
	"fmt"
)

func FindAll() (Streets, error) {
	result := Streets{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+tablename.Street+`" where deleted='false'`)
	return result, err

}
func ExistById(id string) (bool, error) {
	var result bool
	db := connect()
	err := db.Get(&result, `SELECT EXISTS(SELECT * FROM "public"."`+tablename.Street+`" where id='`+id+`')`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func ExistByName(name string) (bool, error) {
	var result bool
	db := connect()
	err := db.Get(&result, `SELECT EXISTS(SELECT * FROM "public"."`+tablename.Street+`" where name='`+name+`')`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneById(id string) (Street, error) {
	result := Street{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+tablename.Street+`" where id='`+id+`'`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneByName(name string) (Street, error) {
	result := Street{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+tablename.Street+`" where name='`+name+`' and deleted='false'`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Street) error {
	db := connect()

	var reqInsert = `INSERT INTO "public"."` + tablename.Street + `" (id, name, type, type_short, city) VALUES (:id, :name, :type, :type_short, :city)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Street) error {
	fmt.Println(entity)
	db := connect()
	var reqUpdate = `UPDATE "public"."` + tablename.Street + `" SET name=:name, type=:type, type_short=:type_short, city=:city, deleted='false' WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveById(id string) {
	db := connect()
	var request = `Update "public"."` + tablename.Street + `" SET deleted='true' WHERE id='` + id + `'`
	db.MustExec(request)
}
