package offers_engine

import "fmt"

type Resource struct {
	Label      string     `json:"label"`
	Name       string     `json:"name"`
	Attributes Attributes `json:"attributes"`
}

type Resources []*Resource

func (r *Resource) Condition(field, operator string, fvalue, value interface{}) (bool, error, interface{}) {
	return r.Attributes.Condition(field, operator, fvalue, value)
}
func (r *Resource) IsOnlyServer(field string) bool {
	return r.Attributes.IsOnlyServer(field)
}
func (rs *Resources) Condition(atr *RuleValue, fvalue interface{}) (bool, error, interface{}) {
	res, err := rs.GetResourceByName(atr.Designator.Resource)
	if err != nil {
		return false, err, nil
	}
	return res.Condition(atr.Designator.Attribute, atr.Operator, fvalue, string(atr.Value))
}
func (rs *Resources) GetResourceByName(name string) (*Resource, error) {
	for _, r := range *rs {
		if r.Name == name {
			return r, nil
		}
	}
	return nil, fmt.Errorf("resource %v not found", name)
}
func (rs *Resources) IsOnlyServer(resource, field string) bool {
	for _, r := range *rs {
		if r.Name == resource {
			return r.IsOnlyServer(field)
		}
	}
	return true
}
