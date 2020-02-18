package mlink

import (
	"github.com/huandu/go-sqlbuilder"
)

func FindAllByOrder(order string) ([]string, error) {
	result := []string{}
	db := connect()
	err := db.Select(&result, `SELECT marker FROM "public"."`+GetTableName()+`" where deleted='false' and "order"=$1`, order)
	return result, err

}
func FindAllByOrderWithDeleted(order string) ([]string, error) {
	result := []string{}
	db := connect()
	err := db.Select(&result, `SELECT marker FROM "public"."`+GetTableName()+`" where "order"=$1`, order)
	return result, err

}
func Insert(marker, order string) error {
	db := connect()

	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					"order", 
					marker 
	) VALUES (
					$1, 
					$2 
					)`
	_, err := db.Queryx(reqInsert, order, marker)
	return err
}
func InsertArray(markers []string, order string) error {
	for _, marker := range markers {
		err := Insert(marker, order)
		if err != nil {
			return err
		}
	}
	return nil
}
func Save(marker, order string) error {
	db := connect()
	var reqUpdate = `UPDATE "public"."` + GetTableName() + `" SET 
					marker=$1, 
	deleted='false' WHERE "order"=$2`
	_, err := db.Queryx(reqUpdate, marker, order)
	return err
}

//сохраняет  маркера заказа
func SaveArray(markers []string, order string) error {
	db := connect()
	sb := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	sb.Update(GetTableName())
	sb.Set("deleted='false'")
	sb.Where(
		sb.In("marker", markersToInterface(markers)...),
		sb.Equal(`"order"`, order),
	)
	sql, args := sb.Build()
	_, err := db.Queryx(sql, args...)
	return err
}

// удаляет все маркеры которых нет в массиве
func RemoveNotExist(markers []string, order string) error {
	db := connect()
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
	_, err := db.Queryx(sql, args...)
	return err
}

//вставить несущесвующие маркеры из массива
func SaveMlinks(markers []string, order string) error {
	if len(markers) == 0 {
		return nil
	}
	db := connect()
	var reqInsert = `INSERT INTO "public"."` + GetTableName() + `" (
					"order", 
					marker 
	) VALUES (
					$1, 
					$2 
					) ON CONFLICT ("order", marker) DO  UPDATE SET deleted='false'`

	tx := db.MustBegin()
	for _, marker := range markers {
		_, err := tx.Exec(reqInsert, order, marker)
		if err != nil {
			return err
		}
	}
	return tx.Commit()

}

func markersToInterface(markers []string) []interface{} {
	result := make([]interface{}, len(markers))
	for i, v := range markers {
		result[i] = v
	}
	return result
}

//Удалить все маркеры заказа
func RemoveByOrder(order string) error {
	db := connect()
	var request = `Update "public"."` + GetTableName() + `" SET deleted='true' WHERE "order"=$1`
	_, err := db.Queryx(request, order)
	return err
}
