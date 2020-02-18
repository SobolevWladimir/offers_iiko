package pmodifiers

import (
	"altegra_offers/lib/base"

	"github.com/huandu/go-sqlbuilder"
)

func FindAll() (Modifiers, error) {
	result := Modifiers{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where deleted='false'`)
	return result, err

}
func FindAllByProduct(product string) (Modifiers, error) {
	result := Modifiers{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where product=$1 and deleted='false'`, product)
	return result, err

}
func FindOneById(id string) (Modifier, error) {
	result := Modifier{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM "public"."`+GetTableName()+`" where id=$1`, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func Insert(entity *Modifier) error {
	db := connect()

	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id, 
					product, 
					modifier, 
					quantity 
	) VALUES (
					:id, 
					:product, 
					:modifier, 
					:quantity 
	)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}
func InsertArray(entitys *Modifiers) error {
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id, 
					product, 
					modifier, 
					quantity 
	) VALUES (
					:id, 
					:product, 
					:modifier, 
					:quantity 
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
func Save(entity *Modifier) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					id=:id, 
					product=:product, 
					modifier=:modifier, 
					quantity=:quantity,  
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
func RemoveNotExist(entitys *Modifiers, product string) error {
	ids := entitys.getIds()
	db := connect()
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update(GetTableName())
	sb.Set("deleted='true'")
	if len(ids) == 0 {
		sb.Where(
			sb.Equal(`"product"`, product),
		)
	} else {
		sb.Where(
			sb.NotIn("id", base.StringToInterface(ids)...),
			sb.Equal(`"product"`, product),
		)

	}
	sql, args := sb.Build()
	_, err := db.Queryx(sql, args...)
	return err
}
func SaveExist(entitys *Modifiers) error {
	if len(*entitys) == 0 {
		return nil
	}
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					id, 
					product, 
					modifier, 
					quantity 
	) VALUES (
					:id, 
					:product, 
					:modifier, 
					:quantity 
					) ON CONFLICT (id) DO  UPDATE SET
					product=:product, 
					modifier=:modifier, 
					quantity=:quantity,  
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
