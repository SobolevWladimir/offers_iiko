package offers_engine

import (
	"altegra_offers/mentity/offerentity"
)

type SetRule struct {
	Description string `json:"description"`
	Effect      bool   `json:"effect"`
	Rules       Rules  `json:"rules"`
}
type SetRules []SetRule

func (r *SetRule) IsOnlyServer() bool {
	return r.Rules.IsOnlyServer()
}
func (rs *SetRules) IsOnlyServer() bool {
	for _, r := range *rs {
		if r.Rules.IsOnlyServer() {
			return true
		}
	}
	return false
}
func (srs *SetRules) Check(given *offerentity.OfferGiven) ([]bool, error, []interface{}) {
	result := []bool{}
	entitys := []interface{}{}
	for _, sr := range *srs {
		res, err, ens := sr.Check(given)
		if err != nil {
			return result, err, nil
		}
		if ens != nil {
			entitys = append(entitys, ens...)
		}
		result = append(result, res)
	}
	return result, nil, entitys
}
func (sr *SetRule) Check(given *offerentity.OfferGiven) (bool, error, []interface{}) {
	ch, err, ent := sr.Rules.Check(given)
	if err != nil {
		return false, err, nil
	}
	if ch {
		return sr.Effect, nil, ent
	}
	return false, nil, ent
}
