package transport

import (
	"encoding/json"
)

type AProductAdded ProductAdded
type AProduct struct {
	AProductItem
	Added AProductAdded `json:"added"`
}

func (a *AProductAdded) UnmarshalJSON(data []byte) error {
	if string(data) == "false" {
		return nil
	}

	entity := ProductAdded{}
	err := json.Unmarshal(data, &entity)
	*a = AProductAdded(entity)
	return err
}
