package coupon

import (
	"altegra_offers/service/coupon"
)

func CheckIsExist(coup *coupon.Coupon) (bool, error) {
	// //все подразделения
	// divs, err := divisions.FindAll()
	// if err != nil {
	//   return false, err
	// }
	// //все категории купонов
	// categorys, err := coupon_category.FindAll()
	// if err != nil {
	//   return false, err
	// }
	// // категория купона
	// coup_cat, err := categorys.GetCategoryById(coup.Category)
	// if err != nil {
	//   return false, err
	// }
	// // подразделение категории купона
	// coup_div, err := divs.GetDivisionById(coup_cat.Division)
	// if err != nil {
	//   return false, err
	// }
	// //все купоны в базе данных
	// cuponsAll, err := coupon.FindAll()
	// if err != nil {
	//   return false, err
	// }
	// divisionAll := divisions.Divisions{}
	// divisionAll = append(divisionAll, coup_div)
	// divisionAll = append(divisionAll, divisions.GetChilds(&coup_div, &divs)...)
	// divisionAll = append(divisionAll, divisions.GetParents(&coup_div, &divs)...)
	// // отфильтрованные категории
	// target_categorys := categorys.FilterByDivision(divisionAll.GetIds())
	//
	// f_coupon := cuponsAll.FilterByCategorys(target_categorys.GetIds())
	// return f_coupon.IsExistByName(coup.Name), nil
	return false, nil
}
func FilterNotExist(coups *coupon.Coupons, category int) (coupon.Coupons, error) {
	result := coupon.Coupons{}
	// //все подразделения
	// divs, err := divisions.FindAll()
	// if err != nil {
	//   return result, err
	// }
	// //все категории купонов
	// categorys, err := coupon_category.FindAll()
	// if err != nil {
	//   return result, err
	// }
	// // категория купона
	// coup_cat, err := categorys.GetCategoryById(category)
	// if err != nil {
	//   return result, err
	// }
	// // подразделение категории купона
	// coup_div, err := divs.GetDivisionById(coup_cat.Division)
	// if err != nil {
	//   return result, err
	// }
	// //все купоны в базе данных
	// cuponsAll, err := coupon.FindAll()
	// if err != nil {
	//   return result, err
	// }
	// divisionAll := divisions.Divisions{}
	// divisionAll = append(divisionAll, coup_div)
	// divisionAll = append(divisionAll, divisions.GetChilds(&coup_div, &divs)...)
	// divisionAll = append(divisionAll, divisions.GetParents(&coup_div, &divs)...)
	// // отфильтрованные категории
	// target_categorys := categorys.FilterByDivision(divisionAll.GetIds())
	//
	// f_coupon := cuponsAll.FilterByCategorys(target_categorys.GetIds())
	// for _, coup := range *coups {
	//   if !f_coupon.IsExistByName(coup.Name) {
	//     result = append(result, coup)
	//   }
	// }
	//
	return result, nil

}
