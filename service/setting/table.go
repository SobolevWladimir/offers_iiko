package setting

func FindAll() (IikoSettings, error) {
	result := IikoSettings{}
	db := connect()
	err := db.Select(&result, `SELECT * FROM `+GetTableName()+` `)
	return result, err
}
