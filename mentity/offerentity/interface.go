package offerentity

import "errors"

type OfferFilterValue struct {
	Field    string
	Operator string
	Value    string
}
type OfferFilterValues []OfferFilterValue
type OfferObjectInterface interface {
	GetValueForOfffer(field string, filters OfferFilterValues) (interface{}, error)
}
type OfferActionInterface interface {
	CalculateClientBonuses(bonuses *ActionBonuses, r *RItems) (float32, error)
}
type OfferGiven struct {
	Order           OfferObjectInterface
	OldVersionOrder OfferObjectInterface
	Client          OfferObjectInterface
	System          OfferObjectInterface
}

func (fs *OfferFilterValues) GetFilterByField(field string) (OfferFilterValue, error) {
	for _, f := range *fs {
		if f.Field == field {
			return f, nil
		}
	}
	return OfferFilterValue{}, errors.New("field not found")
}
