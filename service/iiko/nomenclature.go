package iiko

type Nomenclature struct {
	ProductCategories ProductCategorys `json:"productCategories"`
	Products          Products         `json:"products"`
}

func LoadNomenclature(organization string, auth AuthData) (Nomenclature, error) {
	result := Nomenclature{}

	return result, nil
}
