package offers_engine

import (
	"altegra_offers/service/offers_engine/offer"
)

func FindAllOffers() (Policys, error) {
	result := Policys{}
	offers, err := offer.FindAll()
	if err != nil {
		return result, err
	}
	result = OffersToPolicys(&offers)
	return result, nil
}
func FindOffersByCategory(category int) (Policys, error) {
	result := Policys{}
	offers, err := offer.FindByCategory(category)
	if err != nil {
		return result, err
	}
	result = OffersToPolicys(&offers)
	return result, nil
}
func FindOnePolicyById(id string) (Policy, error) {
	off, err := offer.FindOneById(id)
	if err != nil {
		return Policy{}, err
	}
	return OfferToPolicy(&off), nil
}
func InsertPolicy(entity *Policy) error {
	of, err := entity.toOffer()
	if err != nil {
		return err
	}
	return offer.Insert(&of)
}
func SavePolicy(entity *Policy) error {

	of, err := entity.toOffer()
	if err != nil {
		return err
	}
	return offer.Save(&of)
}
func RemovePolicyById(id int) {
	offer.RemoveById(id)
}
