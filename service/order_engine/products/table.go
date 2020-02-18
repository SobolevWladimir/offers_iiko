package products

import (
	"altegra_offers/lib/base"

	"github.com/huandu/go-sqlbuilder"
)

func FindAll() (Products, error) {
	result := Products{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where deleted='false'`)
	return result, err

}
func FindAllByOrder(order string) (Products, error) {
	result := Products{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where "order"=$1 and deleted='false'`, order)
	return result, err

}
func FindOneById(id string) (Product, error) {
	result := Product{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where id=$1`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Product) error {
	db := connect()

	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id,
					"order", 
					product, 
					quantity, 
					comment
	) VALUES (
					:id,
					:order, 
					:product, 
					:quantity, 
					:comment
					)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func InsertArray(entitys *Products) error {
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id,
					"order", 
					product, 
					quantity, 
					comment
	) VALUES (
					:id,
					:order, 
					:product, 
					:quantity, 
					:comment
					)`
	tx := db.MustBegin()
	for _, entity := range *entitys {
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		if _, err := tx.NamedExec(reqInsert, entity); err != nil {
			return err
		}
	}
	return tx.Commit()
}
func Save(entity *Product) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					"order"=:order, 
					product=:product, 
					quantity=:quantity, 
					comment=:comment,
	deleted='false' WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func RemoveNotExist(entitys *Products, order string) error {
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
func SaveExist(entitys *Products) error {
	if len(*entitys) == 0 {
		return nil
	}
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id,
					"order", 
					product, 
					quantity, 
					comment
	) VALUES (
					:id,
					:order, 
					:product, 
					:quantity, 
					:comment
					) ON CONFLICT (id) DO  UPDATE SET
					"order"=:order, 
					product=:product, 
					quantity=:quantity, 
					comment=:comment,
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
func RemoveById(id string) error {
	db := connect()
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id=$1`
	_, err := db.Queryx(request, id)
	return err
}
