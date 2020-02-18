package client_category_link

func FindAll() (ClientCategoryLinks, error) {
	result := ClientCategoryLinks{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM "public"."`+GetTableName()+`" where deleted='false'`)
	return result, err
}
func FindAllCategoriesByClient(client string) ([]string, error) {
	result := []string{}
	db := connect()
	err := db.Select(&result, `SELECT category FROM "public"."`+GetTableName()+`" where client='`+client+`' and deleted='false'`)
	return result, err
}

func Save(client, category string) error {
	sql := `UPDATE "public"."` + GetTableName() +
		`"SET	deleted='false' WHERE category=$1 and client=$2`
	db := connect()
	_, err := db.Query(sql, category, client)
	return err
}

func Remove(client, category string) error {
	sql := `UPDATE "public"."` + GetTableName() +
		`"SET	deleted='true' WHERE category=$1 and client=$2`
	db := connect()
	_, err := db.Query(sql, category, client)
	return err
}
func Insert(client, category string) error {
	db := connect()
	var sql = `INSERT INTO "public"."` + GetTableName() + `" 
		(client, category)	VALUES
		 ($1, $2)`
	_, err := db.Query(sql, client, category)
	return err
}
func InsertByClient(client string, categories []string) error {
	for _, cat := range categories {
		if err := Insert(client, cat); err != nil {
			return err
		}
	}
	return nil
}
func SaveByClient(client string, categories []string) error {
	// update deleted
	for _, cat := range categories {
		if err := Save(client, cat); err != nil {
			return err
		}
	}
	oldcategs, err := FindAllCategoriesByClient(client)
	if err != nil {
		return err
	}
	added := getNonExist(categories, oldcategs)
	del := getNonExist(oldcategs, categories)
	for _, cat := range added {
		if err := Insert(client, cat); err != nil {
			return err
		}

	}
	for _, cat := range del {
		if err := Remove(client, cat); err != nil {
			return err
		}
	}
	return nil

}
func getNonExist(target []string, source []string) []string {
	result := []string{}
	for _, t := range target {
		if !exist(t, source) {
			result = append(result, t)
		}
	}
	return result
}
func exist(target string, source []string) bool {
	for _, s := range source {
		if s == target {
			return true
		}
	}
	return false
}
