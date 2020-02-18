package offer

import (
	"altegra_offers/service/offers_engine"
	"errors"
)

func FilterByCategoryIn(value interface{}, entitys *offers_engine.Policys) (offers_engine.Policys, error) {
	result := offers_engine.Policys{}

	var id_category int
	var err error
	switch x := value.(type) {
	case int:
		id_category = x
	default:
		return result, errors.New("FilterByCategoryIn:I can not determine the type")
	}
	categAll, err := offers_engine.FindAllCategory()
	if err != nil {
		return result, err
	}
	target_c, err := categAll.GetCategoryById(id_category)
	if err != nil {
		return result, err
	}
	categ_chaild := categAll.GetChilds(&target_c)
	categ_chaild.Append(&target_c)
	return entitys.GetByCategorys(categ_chaild.GetIds()), err
}
func FilterByCategoryEqual(value interface{}, entitys *offers_engine.Policys) (offers_engine.Policys, error) {
	result := offers_engine.Policys{}

	var id int
	switch x := value.(type) {
	case int:
		id = x
	default:
		return result, errors.New("FilterByCategoryIn:I can not determine the type")
	}

	return entitys.GetByCategory(id), nil
}
