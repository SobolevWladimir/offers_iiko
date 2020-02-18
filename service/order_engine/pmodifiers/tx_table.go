package pmodifiers

import (
	"altegra_offers/lib/base"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func TxInsert(tx *sqlx.Tx, entity *Modifier) error {
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
	_, err := tx.NamedExec(reqInsert, entity)
	return err
}
func TxInsertArray(tx *sqlx.Tx, entitys *Modifiers) error {
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
	for _, entity := range *entitys {
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		if _, err := tx.NamedExec(reqInsert, entity); err != nil {
			return err
		}
	}
	return nil
}
func TxSave(tx *sqlx.Tx, entity *Modifier) error {
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					id=:id, 
					product=:product, 
					modifier=:modifier, 
					quantity=:quantity,  
	deleted='false' WHERE id=:id`
	_, err := tx.NamedExec(reqUpdate, entity)
	return err
}
func TxRemoveById(tx *sqlx.Tx, id string) error {
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id=$1`
	_, err := tx.Queryx(request, id)
	return err
}
func TxRemoveNotExist(tx *sqlx.Tx, entitys *Modifiers, product string) error {
	ids := entitys.getIds()
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
	_, err := tx.Exec(sql, args...)
	return err
}
func TxSaveExist(tx *sqlx.Tx, entitys *Modifiers) error {
	if len(*entitys) == 0 {
		return nil
	}
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

	for _, entity := range *entitys {
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		_, err := tx.NamedExec(reqInsert, entity)
		if err != nil {
			return err
		}
	}
	return nil
}
