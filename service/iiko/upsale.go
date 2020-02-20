package iiko

import "offers_iiko/mentity/offerentity"

type Upsale struct {
	SourceActionId string `json:"sourceActionId"` // Id действия, добавившего предложение
	SuggestionText string `json:"suggestionText"` // Текст подсказки оператору

}
type Upsales []Upsale

func (up *Upsales) GetActons() offerentity.Actions {
	result := offerentity.Actions{}
	for _, sal := range *up {
		result = append(result, offerentity.Action{
			Type: offerentity.TypeUpsale,
			Data: sal,
		})
	}
	return result
}
