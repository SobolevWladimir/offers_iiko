package iiko

type FreeProductsGroup struct {
	SourceActionId     string   `json:"sourceActionId"`
	DescriptionForUser string   `json:"descriptionForUser"`
	ProductCodes       []string `json:"productCodes"`
}
type FreeProductsGroups []FreeProductsGroup
