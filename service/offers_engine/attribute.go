package offers_engine

import "fmt"

type ConditionFn func(fvalue, value interface{}) (bool, error, interface{})
type Operator struct {
	Label     string      `json:"label"`
	Name      string      `json:"name"`
	Input     string      `json:"input"`
	TypeV     string      `json:"type"`
	Condition ConditionFn `json:"-"`
}
type Operators []Operator

type AttributeFilter struct {
	Label     string    `json:"label"`
	Name      string    `json:"name"`
	Operators Operators `json:"operators"`
}
type AttributeFilters []AttributeFilter

type Attribute struct {
	Label      string           `json:"label"`
	Name       string           `json:"name"`
	ServerOnly bool             `json:"server_only"`
	Operators  Operators        `json:"operators"`
	Filters    AttributeFilters `json:"filters"`
}
type Attributes []Attribute

func (os *Operators) Condition(operator string, fvalue, value interface{}) (bool, error, interface{}) {
	for _, o := range *os {
		if o.Name == operator {
			if o.Condition == nil {
				return false, fmt.Errorf("operator (%v): condition is null", operator), nil
			}
			return o.Condition(fvalue, value)
		}

	}
	return false, fmt.Errorf("operator (%v) not found ", operator), nil
}

func (a *Attribute) Condition(operator string, fvalue, value interface{}) (bool, error, interface{}) {
	return a.Operators.Condition(operator, fvalue, value)
}
func (as *Attributes) Condition(field, operator string, fvalue, value interface{}) (bool, error, interface{}) {
	for _, a := range *as {
		if a.Name == field {
			return a.Condition(operator, fvalue, value)
		}
	}
	return false, fmt.Errorf(" field %v not found", field), nil
}
func (as *Attributes) IsOnlyServer(field string) bool {
	for _, a := range *as {
		if a.Name == field {
			return a.ServerOnly
		}

	}
	return true
}
