package offerentity

import (
	"altegra_offers/lib/log"
	"encoding/json"
)

type ActionData string

func (a *ActionData) ToActinBonuses() ActionBonuses {
	result := ActionBonuses{}
	data := []byte(*a)
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.MEntityError(err)
	}
	return result
}
func (a *ActionData) ToActinCoupon() ActionCoupon {
	result := ActionCoupon{}
	data := []byte(*a)
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.MEntityError(err)
	}
	return result
}
func (a *ActionData) ToActionSpecialDiscount() ActionSpecialDiscount {
	result := ActionSpecialDiscount{}
	data := []byte(*a)
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.MEntityError(err)
	}
	return result
}
func (a *ActionData) ToActionDiscount() ActionDiscount {
	result := ActionDiscount{}
	data := []byte(*a)
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.MEntityError(err)
	}
	return result
}
func (a *ActionData) ToActionClientCategory() ActionClientCategory {
	result := ActionClientCategory{}
	data := []byte(*a)
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.MEntityError(err)
	}
	return result
}
func (a *ActionData) UnmarshalJSON(data []byte) error {
	*a = ActionData(data)
	return nil
}

func (a ActionData) MarshalJSON() ([]byte, error) {
	return []byte(a), nil
}
