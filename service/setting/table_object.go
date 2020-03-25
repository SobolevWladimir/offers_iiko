package setting

type TableObject struct {
}

func (t *TableObject) FindCityNameById(id int) (string, error) {
	var result string
	db := connect()
	err := db.Get(&result, `SELECT name FROM sity WHERE id=?`, id)
	return result, err
}
