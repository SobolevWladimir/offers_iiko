package offers_engine

import (
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/offers_engine/offer"
)

func (entitys *Policys) toOffers() (offer.Offers, error) {
	result := offer.Offers{}
	for _, entity := range *entitys {
		of, err := entity.toOffer()
		if err != nil {
			return result, err
		}
		result = append(result, of)
	}
	return result, nil
}
func (p *Policys) GetByCategorys(ids []int) Policys {
	result := Policys{}
	for _, id := range ids {
		result = append(result, p.GetByCategory(id)...)
	}
	return result
}
func (p *Policys) GetByCategory(id int) Policys {
	result := Policys{}
	for _, pl := range *p {
		if pl.Category == id {
			result = append(result, pl)
		}
	}
	return result
}
func (p *Policys) FilterByStatus(status string) Policys {
	result := Policys{}
	for _, pol := range *p {
		if pol.Status == status {
			result = append(result, pol)
		}
	}
	return result
}
func (ps *Policys) IsOnlyServer() bool {
	for _, p := range *ps {
		if p.IsOnlyServer() {
			return true
		}
	}
	return false
}

//калькулирует акции для заказа
func (p *Policys) Calculate(given *offerentity.OfferGiven) (offerentity.Offers, error) {
	result := offerentity.Offers{}
	for _, pol := range *p {
		if pol.Active == false {
			continue
		}
		item, err := pol.Calculate(given)
		if err != nil {
			return result, err
		}
		result = append(result, item)
	}

	return result, nil
}
