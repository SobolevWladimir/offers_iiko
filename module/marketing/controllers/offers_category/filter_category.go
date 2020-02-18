package offers_category

import (
	"altegra_offers/service/offers_engine"
	"errors"
)

func FilterByCategoryIn(value interface{}, entitys *offers_engine.Categorys) (offers_engine.Categorys, error) {
	result := offers_engine.Categorys{}
	var err error

	var id_category int
	switch x := value.(type) {
	case int:
		id_category = x
	default:
		return result, errors.New("FilterByCategoryIn:I can not determine the type")
	}
	target_div, err := entitys.GetCategoryById(id_category)

	divs_child := entitys.GetChilds(&target_div)
	result.Append(&target_div)
	result.AppendCategorys(&divs_child)
	return result, err
}
func FilterByCategoryEqual(value interface{}, entitys *offers_engine.Categorys) (offers_engine.Categorys, error) {
	result := offers_engine.Categorys{}
	var err error

	var id int
	switch x := value.(type) {
	case int:
		id = x
	default:
		return result, errors.New("FilterByCategoryIn:I can not determine the type")
	}
	target, err := entitys.GetCategoryById(id)
	if err != nil {
		return result, nil
	}
	target.Parent.Valid = false
	result.Append(&target)
	return result, err
}
