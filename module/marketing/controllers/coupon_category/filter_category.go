package coupon_category

import (
	"altegra_offers/service/coupon_category"
	"errors"
)

func FilterByCategoryIn(value interface{}, entitys *coupon_category.Categorys) (coupon_category.Categorys, error) {
	result := coupon_category.Categorys{}
	var id_category int
	var err error
	switch x := value.(type) {
	case int:
		id_category = x
	default:
		return result, errors.New("FilterByCategoryIn:I can not determine the type")
	}
	divsall, err := coupon_category.FindAll()
	if err != nil {
		return result, err
	}
	target_div, err := divsall.GetCategoryById(id_category)

	if err != nil {
		return result, err
	}
	divs_child := coupon_category.GetChilds(&target_div, &divsall)
	result = append(divs_child, target_div)
	return result, err
}
func FilterByCategoryEqual(value interface{}, entitys *coupon_category.Categorys) (coupon_category.Categorys, error) {
	result := coupon_category.Categorys{}
	var id int
	var err error
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
	result = entitys.GetCategorysByParent(int64(id))
	result = append(result, target)
	return result, err
}
