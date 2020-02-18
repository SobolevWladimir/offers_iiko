package mlink

import (
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

func TxInsert(tx *sqlx.Tx, marker, order string) error {

	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					"order", 
					marker 
	) VALUES (
					$1, 
					$2 
					)`
	_, err := tx.Queryx(reqInsert, order, marker)
	return err
}
func TxInsertArray(tx *sqlx.Tx, markers []string, order string) error {
	for _, marker := range markers {
		err := Insert(marker, order)
		if err != nil {
			return err
		}
	}
	return nil
}
func TxSave(tx *sqlx.Tx, marker, order string) error {
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					marker=$1, 
	deleted='false' WHERE "order"=$2`
	_, err := tx.Queryx(reqUpdate, marker, order)
	return err
}

//сохраняет  маркера заказа
func TxSaveArray(tx *sqlx.Tx, markers []string, order string) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update(GetTableName())
	sb.Set("deleted='false'")
	sb.Where(
		sb.In("marker", markersToInterface(markers)...),
		sb.Equal(`"order"`, order),
	)
	sql, args := sb.Build()
	_, err := tx.Queryx(sql, args...)
	return err
}

// удаляет все маркеры которых нет в массиве
func TxRemoveNotExist(tx *sqlx.Tx, markers []string, order string) error {
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update(GetTableName())
	sb.Set("deleted='true'")
	if len(markers) == 0 {
		sb.Where(
			sb.Equal(`"order"`, order),
		)
	} else {
		sb.Where(
			sb.NotIn("marker", markersToInterface(markers)...),
			sb.Equal(`"order"`, order),
		)

	}

	sql, args := sb.Build()
	_, err := tx.Exec(sql, args...)
	return err
}

//вставить несущесвующие маркеры из массива
func TxSaveMlinks(tx *sqlx.Tx, markers []string, order string) error {
	if len(markers) == 0 {
		return nil
	}
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					"order", 
					"marker" 
	) VALUES (
					$1, 
					$2 
					) ON CONFLICT ("order", "marker") DO  UPDATE SET deleted='false'`

	for _, marker := range markers {
		_, err := tx.Exec(reqInsert, order, marker)
		if err != nil {
			return err
		}
	}
	return nil
}

//Удалить все маркеры заказа
func TxRemoveByOrder(tx *sqlx.Tx, order string) error {
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE "order"=$1`
	_, err := tx.Queryx(request, order)
	return err
}
