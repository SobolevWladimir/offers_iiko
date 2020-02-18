package coupon

import (
	"gopkg.in/guregu/null.v3"
)

type Coupon struct {
	Id       int         `db:"id" json:"id" valid:"-"`
	Name     string      `db:"name" json:"name" valid:"required" accessfield:"name"`
	Status   bool        `db:"status" json:"status" valid:"-"`
	Type     int         `db:"type" json:"type" valid:"-"`
	Comment  null.String `db:"comment" json:"comment" valid:"-"`
	Category int         `db:"category" json:"category" valid:"-"`
	Sort     int         `db:"sort" json:"sort" valid:"-"`
	Deleted  bool        `db:"deleted" json:"-" valid:"-"`
}
type Coupons []Coupon

const (
	CouponTypeOne   = 0
	CouponTypeMulti = 1
)

//отфильторовать купоны по категориям
func (coups *Coupons) FilterByCategorys(categorys []int) Coupons {
	result := Coupons{}
	for _, category := range categorys {
		result = append(result, coups.GetCouponsByCategory(category)...)
	}
	return result
}

//отфильторовать купоны по  стутусу
func (coups *Coupons) FilterByStatus(status bool) Coupons {
	result := Coupons{}
	for _, c := range *coups {
		if c.Status == status {
			result = append(result, c)
		}
	}
	return result
}

//Найти купон по категории
func (coups *Coupons) GetCouponsByCategory(category int) Coupons {
	result := Coupons{}
	for _, coup := range *coups {
		if coup.Category == category {
			result = append(result, coup)
		}
	}
	return result
}
func (coups *Coupons) GetNames() []string {
	result := []string{}
	for _, c := range *coups {
		result = append(result, c.Name)
	}
	return result
}
func (coups *Coupons) GetByCategorys(categorys []int) Coupons {
	result := Coupons{}
	for _, cat := range categorys {
		result = append(result, coups.GetCouponsByCategory(cat)...)
	}
	return result
}

//Проверяет по имени  наличие купона в списке
func (coups *Coupons) IsExistByName(name string) bool {
	for _, coup := range *coups {
		if coup.Name == name {
			return true
		}
	}
	return false
}
