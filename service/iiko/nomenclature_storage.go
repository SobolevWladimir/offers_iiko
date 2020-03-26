package iiko

var nomeclatureStorage map[int]Nomenclature = make(map[int]Nomenclature)

//Обновить номенклатуру для городов
func UpdataStorageNomenclature(data map[int]AuthData) error {
	for i, auth := range data {
		nom, err := LoadNomenclature(auth)
		if err != nil {
			return err
		}
		nomeclatureStorage[i] = nom
	}
	return nil
}
func StorageGenNomenclatureByKey(key int) Nomenclature {
	return nomeclatureStorage[key]

}
