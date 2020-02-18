package coupon

import (
	"altegra_offers/service/coupon"
	"altegra_offers/service/coupon_category"
	"errors"
)

func FilterByCategoryIn(value interface{}, entitys *coupon.Coupons) (coupon.Coupons, error) {
	result := coupon.Coupons{}

	var id_category int
	var err error
	switch x := value.(type) {
	case int:
		id_category = x
	default:
		return result, errors.New("FilterByCategoryIn:I can not determine the type")
	}
	categAll, err := coupon_category.FindAll()
	if err != nil {
		return result, err
	}
	target_c, err := categAll.GetCategoryById(id_category)
	if err != nil {
		return result, err
	}
	categ_chaild := coupon_category.GetChilds(&target_c, &categAll)
	categs := append(categ_chaild, target_c)
	return entitys.GetByCategorys(categs.GetIds()), err
}
func FilterByCategoryEqual(value interface{}, entitys *coupon.Coupons) (coupon.Coupons, error) {
	result := coupon.Coupons{}

	var id int
	switch x := value.(type) {
	case int:
		id = x
	default:
		return result, errors.New("FilterByCategoryIn:I can not determine the type")
	}

	return entitys.GetCouponsByCategory(id), nil
}
