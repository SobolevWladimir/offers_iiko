package products

import (
	"altegra_offers/lib/base"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func TxInsert(tx *sqlx.Tx, entity *Product) error {
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
	_, err := tx.NamedExec(reqInsert, entity)
	return err
}
func TxInsertArray(tx *sqlx.Tx, entitys *Products) error {
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
func TxSave(tx *sqlx.Tx, entity *Product) error {
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					"order"=:order, 
					product=:product, 
					quantity=:quantity, 
					comment=:comment,
	deleted='false' WHERE id=:id`
	_, err := tx.NamedExec(reqUpdate, entity)
	return err
}
func TxRemoveNotExist(tx *sqlx.Tx, entitys *Products, order string) error {
	ids := entitys.getIds()
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
	_, err := tx.Exec(sql, args...)
	return err
}
func TxSaveExist(tx *sqlx.Tx, entitys *Products) error {
	if len(*entitys) == 0 {
		return nil
	}
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
func TxRemoveById(tx *sqlx.Tx, id string) error {
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id=$1`
	_, err := tx.Queryx(request, id)
	return err
}
