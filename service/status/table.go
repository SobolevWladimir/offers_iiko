package status

import "altegra_offers/lib/tablename"

func FindAll() (Statuses, error) {
	result := Statuses{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+tablename.Status+`" where deleted='false' ORDER BY priority`)
	return result, err

}

func ExistByName(name string) (bool, error) {
	var result bool
	db := connect()
	err := db.Get(&result, `SELECT EXISTS(SELECT * FROM "public"."`+tablename.Status+`" where name='`+name+`')`)
	if err != nil {
		return result, err
	}
	println(result)
	return result, nil
}

func FindOneByName(name string) (Status, error) {
	result := Status{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+tablename.Status+`" where name='`+name+`'`)
	if err != nil {
		return result, err
	}
	return result, nil
}
func FindOneByCode(code string) (Status, error) {
	result := Status{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+tablename.Status+`" where code='`+code+`'`)
	if err != nil {
		return result, err
	}
	return result, nil
}

func Insert(entity *Status) error {
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + tablename.Status + `" ( name, code, priority,comment,color) VALUES ( :name, :code,  :priority, :comment, :color )`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Status) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + tablename.Status + `" SET  code=:code, priority=:priority, comment=:comment, color=:color, deleted='false' WHERE name=:name`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveByCode(code string) {
	db := connect()
	var request = `Update "public"."` + tablename.Status + `" SET deleted='true' WHERE name='` + code + `'`
	db.MustExec(request)
}
