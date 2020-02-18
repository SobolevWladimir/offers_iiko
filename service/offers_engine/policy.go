package offers_engine

import (
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/coupon"
	"altegra_offers/service/offers_engine/offer"
	"encoding/json"

	"gopkg.in/guregu/null.v3"
)

type Policy struct {
	Id        int      `json:"id" valid:"-"`
	Active    bool     `json:"active" valid:"-"`
	Name      string   `json:"name" valid:"-"`
	Status    string   `json:"status" valid:"-"`
	Algorithm int      `json:"combining" valid:"-"`
	SetRules  SetRules `json:"setrules" valid:"-"`
	Actions   Actions  `json:"actions" valid:"-"`
	Category  int      `json:"category" valid:"-"`
	Sort      int      `json:"sort" valid:"-"`
}
type Policys []Policy

func OfferToPolicy(entity *offer.Offer) Policy {
	result := Policy{}
	result.Id = entity.Id
	result.Active = entity.Active
	result.Name = entity.Name
	result.Status = entity.Status
	result.Algorithm = entity.Algorithm
	result.Category = entity.Category
	result.Sort = entity.Sort
	var rules SetRules
	json.Unmarshal([]byte(entity.SetRules.String), &rules)
	actions := Actions{}
	json.Unmarshal([]byte(entity.Actions.String), &actions)
	result.SetRules = rules
	result.Actions = actions
	return result
}
func OffersToPolicys(entitys *offer.Offers) Policys {
	result := Policys{}
	for _, entity := range *entitys {
		result = append(result, OfferToPolicy(&entity))
	}
	return result
}
func (entity *Policy) toOffer() (offer.Offer, error) {
	result := offer.Offer{}
	result.Id = entity.Id
	result.Active = entity.Active
	result.Name = entity.Name
	result.Status = entity.Status
	result.Algorithm = entity.Algorithm
	result.Category = entity.Category
	result.Sort = entity.Sort
	rules, err := json.Marshal(entity.SetRules)
	if err != nil {
		return result, err
	}
	result.SetRules = null.StringFrom(string(rules))
	actions, err := json.Marshal(entity.Actions)
	if err != nil {
		return result, err
	}
	result.Actions = null.StringFrom(string(actions))
	return result, nil
}
func (p *Policy) IsOnlyServer() bool {
	return p.SetRules.IsOnlyServer()
}

func (p *Policy) Calculate(given *offerentity.OfferGiven) (offerentity.Offer, error) {
	result := offerentity.Offer{}
	result.Id = p.Id

	//находим положительные
	iswork, err, ents := p.IsWorked(given)
	if err != nil {
		result.Error = true
		result.Message = err.Error()
		result.Actions = offerentity.Actions{}
		return result, nil
	}
	if iswork {
		result.Actions = p.Actions.ToOfferAction()
		result.Actions = append(result.Actions, p.GetSpecAction(ents)...)
	} else {
		result.Message = "negative counting result"
		result.Actions = offerentity.Actions{}
	}

	return result, nil
}
func (p *Policy) GetSpecAction(entitys []interface{}) offerentity.Actions {
	result := offerentity.Actions{}
	for _, ent := range entitys {
		switch x := ent.(type) {
		case coupon.Coupon:
			if x.Type == coupon.CouponTypeOne {
				byt, err := json.Marshal(x)
				if err != nil {
					continue
				}
				result = append(result, offerentity.Action{
					IsDone: 0,
					Type:   offerentity.TypeActionCoupon,
					Data:   offerentity.ActionData(byt),
				})
			}
		default:

		}

	}

	return result
}
func (p *Policy) IsWorked(given *offerentity.OfferGiven) (bool, error, []interface{}) {

	effects, err, ents := p.SetRules.Check(given)
	if err != nil {
		return false, err, nil
	}

	trueffec := []bool{}
	for _, ef := range effects {
		if ef {
			trueffec = append(trueffec, ef)
		}
	}
	var count int
	if p.Algorithm > 0 {
		count = p.Algorithm
	} else {
		count = len(effects)
	}
	if count == 0 {
		return false, nil, nil
	}
	return count <= len(trueffec), nil, ents
}
