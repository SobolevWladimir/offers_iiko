package iiko

type AvailableCombo struct {
	SpecificationId string              `json:"specificationId"`
	GroupMapping    []ComboGroupMapping `json:"groupMapping"`
}

type AvailableCombos []AvailableCombo
type ComboGroupMapping struct {
	GroupId string `json:"groupId"`
	ItemId  string `json:"itemId"`
}
