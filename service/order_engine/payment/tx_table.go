package payment

import (
	"altegra_offers/lib/base"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func TxInsert(tx *sqlx.Tx, entity *Payment) error {
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
	_, err := tx.NamedExec(reqInsert, entity)
	return err
}
func TxSave(tx *sqlx.Tx, entity *Payment) error {
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					type=:type,
					amount=:amount,
					deleted='false' WHERE "order"=:order`
	_, err := tx.NamedExec(reqUpdate, entity)
	return err
}
func TxRemoveById(tx *sqlx.Tx, id string) error {
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE id=$1`
	_, err := tx.Queryx(request, id)
	return err
}
func TxInsertArray(tx *sqlx.Tx, entitys *Payments) error {
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
	for _, entity := range *entitys {
		if len(entity.Id) == 0 {
			entity.Id = base.UUID()
		}
		_, err := tx.NamedExec(reqInsert, &entity)
		if err != nil {
			return err
		}
	}
	return nil
}
func TxRemoveNotExist(tx *sqlx.Tx, entitys *Payments, order string) error {
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
func TxSaveExist(tx *sqlx.Tx, entitys *Payments) error {
	if len(*entitys) == 0 {
		return nil
	}
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
