package payment

import (
	"altegra_offers/lib/base"

	"github.com/huandu/go-sqlbuilder"
)

func FindAll() (Payments, error) {
	result := Payments{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where deleted='false'`)
	return result, err

}
func FindAllByOrder(order string) (Payments, error) {
	result := Payments{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where "order"=$1 and deleted='false'`, order)
	return result, err

}
func FindOneByOrder(name string) (Payment, error) {
	result := Payment{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where order=$1 and deleted='false'`, name)
	if err != nil {
		return result, err
	}
	return result, nil
}

func Insert(entity *Payment) error {
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id,
					"order",
					type,
					amount,
					) VALUES (
					:id,
					:order,
					:type,
					:amount
	)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func Save(entity *Payment) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					type=:type,
					amount=:amount,
					deleted='false' WHERE "order"=:order`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveById(id string) error {
	db := connect()
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id=$1`
	_, err := db.Queryx(request, id)
	return err
}
func InsertArray(entitys *Payments) error {
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id,
					"order",
					type,
					amount
					) VALUES (
					:id,
					:order,
					:type,
					:amount
	)`
	tx := db.MustBegin()
	for _, entity := range *entitys {
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		_, err := tx.NamedExec(reqInsert, &entity)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
func RemoveNotExist(entitys *Payments, order string) error {
	ids := entitys.getIds()
	db := connect()
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update(GetTableName())
	sb.Set("deleted='true'")
	if len(ids) == 0 {
		sb.Where(
			sb.Equal(`"order"`, order),
		)
	} else {
		sb.Where(
			sb.NotIn("id", base.StringToInterface(ids)...),
			sb.Equal(`"order"`, order),
		)

	}
	sql, args := sb.Build()
	_, err := db.Queryx(sql, args...)
	return err
}
func SaveExist(entitys *Payments) error {
	if len(*entitys) == 0 {
		return nil
	}
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id, 
					"order", 
					type, 
					amount
	) VALUES (
					:id,
					:order,
					:type,
					:amount
					) ON CONFLICT (id) DO  UPDATE SET
					type=:type,
					amount=:amount,
					deleted='false'`

	tx := db.MustBegin()
	for _, entity := range *entitys {
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		_, err := tx.NamedExec(reqInsert, entity)
		if err != nil {
			return err
		}
	}
	return tx.Commit()

}
