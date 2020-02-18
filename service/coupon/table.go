package coupon

import (
	"altegra_offers/lib/base"

	"github.com/huandu/go-sqlbuilder"
)

//поиск всеx купонов
func FindAll() (Coupons, error) {
	result := Coupons{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM `+GetTableName()+` `)
	return result, err

}
func FindByCategory(category int) (Coupons, error) {
	result := Coupons{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM `+GetTableName()+` where category=?`, category)
	return result, err

}

//Поиск купонов в по категориям
func FindByCategorys(cats []int) (Coupons, error) {
	result := Coupons{}
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	sb.Select("*")
	sb.From(GetTableName())
	sb.Where(
		sb.In("category", base.IntToInterface(cats)...),
		sb.Equal("deleted", "false"),
	)
	db := connect()
	sql, args := sb.Build()
	err := db.Select(&result, sql, args...)
	return result, err
}
func FindByName(name string) (Coupons, error) {
	result := Coupons{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where name=? `, name)
	if err != nil {
		return result, err
	}
	return result, nil
}

//Поис купона по имени
func FindOneByName(name string) (Coupon, error) {
	result := Coupon{}
	db := connect()
	err := db.Get(&result, `SELECT * FROM `+GetTableName()+` where name=? `, name)
	if err != nil {
		return result, err
	}
	return result, nil
}

//Вставить купон
func Insert(entity *Coupon) error {
	db := connect()

	var reqInsert = `INSERT INTO ` + GetTableName() + `
	(id, name, status, type, comment, sort, category) VALUES 
	(:id, :name, :status, :type, :comment, :sort, :category)`
	_, err := db.NamedExec(reqInsert, entity)
	return err
}

//Вставить купоны
func InsertCoupons(entitys *Coupons) error {
	db := connect()

	var reqInsert = `INSERT INTO ` + GetTableName() + `
	(id, name, status, type, comment, sort, category) VALUES 
	(:id, :name, :status, :type, :comment, :sort, :category)`
	tx := db.MustBegin()
	for _, entity := range *entitys {
		_, err := tx.NamedExec(reqInsert, entity)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

//Сохранить купоны
func Save(entity *Coupon) error {
	db := connect()
	var reqUpdate = `UPDATE ` + GetTableName() + ` SET 
	name=:name, 
	status=:status,
	type=:type,  
	comment=:comment, 
	sort=:sort, 
	category=:category
	WHERE id=:id`
	_, err := db.NamedExec(reqUpdate, entity)
	return err
}
func SaveStasusByIds(ids []string, status bool, comment string) error {
	if len(ids) == 0 {
		return nil
	}
	db := connect()
	tx := db.MustBegin()
	var reqUpdate = `UPDATE ` + GetTableName() + ` SET 
	status=?,
	comment=? 
	WHERE id=?;`
	for _, id := range ids {
		tx.MustExec(reqUpdate, status, comment, id)
	}
	return tx.Commit()

}

//Удалить купон по  id
func RemoveById(id int) {
	db := connect()
	var request = `DELETE FROM  ` + GetTableName() + ` WHERE id=?`
	db.MustExec(request, id)
}
