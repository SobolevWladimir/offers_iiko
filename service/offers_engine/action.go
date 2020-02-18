package offers_engine

import (
	"altegra_offers/mentity/offerentity"
	"encoding/json"
	"fmt"
	"reflect"
)

type Action struct {
	Id    string `json:"id"`
	Type  int    `json:"type"`
	Value string `json:"value"`
}
type Actions []Action

func (a *Action) ToOfferAction() offerentity.Action {
	result := offerentity.Action{}
	result.Type = a.Type
	result.Data = offerentity.ActionData(a.Value)
	return result
}
func (as *Actions) ToOfferAction() offerentity.Actions {
	result := offerentity.Actions{}
	for _, a := range *as {
		result = append(result, a.ToOfferAction())
	}
	return result
}

func (a *Action) UnmarshalJSON(data []byte) error {
	var err error
	raw := make(map[string]interface{})
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if v, contains := raw["type"]; contains {
		switch x := v.(type) {
		case int:
			a.Type = x
		case float64:
			a.Type = int(x)
		default:
			err = fmt.Errorf("json: cannot unmarshal %v into Go value of type int", reflect.TypeOf(v).Name())
		}
	}
	if v, contains := raw["value"]; contains {
		var str []byte
		str, err = json.Marshal(v)
		a.Value = string(str)

	}
	return err
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this String is null.
func (a Action) MarshalJSON() ([]byte, error) {
	raw := make(map[string]interface{})
	raw["type"] = a.Type
	var value interface{}
	if err := json.Unmarshal([]byte(a.Value), &value); err != nil {
		return []byte{}, err
	}
	raw["value"] = value
	return json.Marshal(raw)
}
