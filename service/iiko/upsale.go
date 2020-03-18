package iiko

import (
	"encoding/json"
	"offers_iiko/mentity/offerentity"
)

type Upsale struct {
	SourceActionId string `json:"sourceActionId"` // Id действия, добавившего предложение
	SuggestionText string `json:"suggestionText"` // Текст подсказки оператору
}
type OfferInfo struct {
	Type         int    `json:"type"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IikoActionId string `json:"iiko_action_id"`
	Target       string `json:"target"`
}
type Upsales []Upsale

func (up *Upsales) GetActons() offerentity.Actions {
	result := offerentity.Actions{}
	for _, sal := range *up {
		offer, err := sal.ToOfferInfo()
		if err != nil {
			continue
		}
		result = append(result, offerentity.Action{
			Type: offerentity.TypeUpsale,
			Data: offer,
		})
	}
	return result
}
func (u *Upsale) ToOfferInfo() (OfferInfo, error) {
	result := OfferInfo{}

	err := json.Unmarshal([]byte(u.SuggestionText), &result)
	return result, err
}
