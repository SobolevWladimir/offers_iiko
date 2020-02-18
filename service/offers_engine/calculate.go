package offers_engine

import (
	"altegra_offers/mentity/offerentity"
	"errors"
)

//калькулирует акции для заказа, но не выполняет действия
func Calculate(given *offerentity.OfferGiven, tpolicys Policys) (offerentity.Offers, error) {
	result := offerentity.Offers{}
	if !AllowCalculte(tpolicys) {
		return result, errors.New("allow calculate on only global server")
	}
	cal, err := tpolicys.Calculate(given)
	if err != nil {
		return result, err
	}
	result = cal
	return result, nil
}

// Проверяет разрешен ли полсчет акций на этой машине
func AllowCalculte(tpolicys Policys) bool {
	if !setting.Client.Run || setting.Client.CalculateOfferInClient {
		return true
	}
	if tpolicys.IsOnlyServer() {
		return false
	}
	return true
}
func GetOffersByOrder(status string, city int) (Policys, error) {
	result := Policys{}
	allPolicys, err := FindAllOffers()
	if err != nil {
		return result, err
	}
	categorys, err := FindByCity(city)
	if err != nil {
		return result, err
	}
	tpolicys := allPolicys.GetByCategorys(categorys.GetIds())
	tpolicys = tpolicys.FilterByStatus(status)
	return tpolicys, nil
}
func GetOfferByCategorys(cats *Categorys) (Policys, error) {
	result := Policys{}
	allPolicys, err := FindAllOffers()
	if err != nil {
		return result, err
	}
	return allPolicys.GetByCategorys(cats.GetIds()), nil

}
