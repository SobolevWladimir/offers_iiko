package offers_engine

import (
	"altegra_offers/lib/base"
	"altegra_offers/mentity/offerentity"
	"fmt"
)

type AttributeDesignator struct {
	Resource  string `json:"resource"`
	Attribute string `json:"attribute"`
}
type RuleValue struct {
	Designator AttributeDesignator `json:"designator"`
	Operator   string              `json:"operator"`
	Value      base.StringInt      `json:"value"`
}
type RuleFilter struct {
	Attribute string         `json:"attribute"`
	Operator  string         `json:"operator"`
	Value     base.StringInt `json:"value"`
}
type RuleFilters []RuleFilter
type Rule struct {
	Condition RuleValue   `json:"condition"`
	Filters   RuleFilters `json:"filters"`
}
type Rules []Rule

func (rf *RuleFilter) ToFilterValue() offerentity.OfferFilterValue {
	return offerentity.OfferFilterValue{
		Field:    rf.Attribute,
		Operator: rf.Operator,
		Value:    string(rf.Value),
	}
}
func (rfs *RuleFilters) ToFilterValue() offerentity.OfferFilterValues {
	result := offerentity.OfferFilterValues{}
	for _, rf := range *rfs {
		result = append(result, rf.ToFilterValue())
	}
	return result
}
func (a *AttributeDesignator) IsOnlyServer() bool {
	return StandartResources.IsOnlyServer(a.Resource, a.Attribute)
}

func (r *Rule) IsOnlyServer() bool {
	return r.Condition.Designator.IsOnlyServer()
}
func (rs *Rules) IsOnlyServer() bool {
	for _, r := range *rs {
		if r.IsOnlyServer() {
			return true
		}
	}
	return false
}
func (rs *Rules) Check(given *offerentity.OfferGiven) (bool, error, []interface{}) {
	entitys := []interface{}{}
	for _, r := range *rs {
		res, err, entity := r.Check(given)
		if err != nil {
			return false, err, nil
		}
		if res == false {
			return false, nil, nil
		}
		if entity != nil {
			entitys = append(entitys, entity)
		}

	}
	return true, nil, entitys
}
func (r *Rule) Check(given *offerentity.OfferGiven) (bool, error, interface{}) {
	value, err := r.Condition.Designator.GetValue(given, r.Filters.ToFilterValue())
	if err != nil {
		return false, err, nil
	}
	if value == nil {
		return false, nil, nil
	}
	return StandartResources.Condition(&r.Condition, value)
}
func (at *AttributeDesignator) GetValue(given *offerentity.OfferGiven, filters offerentity.OfferFilterValues) (interface{}, error) {
	switch at.Resource {
	case "client":
		{
			if given.Client == nil {
				return nil, nil
			}
			return given.Client.GetValueForOfffer(at.Attribute, filters)
		}
	case "order":
		return given.Order.GetValueForOfffer(at.Attribute, filters)
	case "system":
		return given.System.GetValueForOfffer(at.Attribute, filters)
	default:
		return nil, fmt.Errorf("not found resource %v", at.Resource)
	}
}
