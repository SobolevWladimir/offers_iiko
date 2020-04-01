package iiko

import (
	"offers_iiko/mentity/offerentity"
	"offers_iiko/mentity/transport"
)

type FreeProductsGroup struct {
	SourceActionId     string   `json:"sourceActionId"`
	DescriptionForUser string   `json:"descriptionForUser"`
	ProductCodes       []string `json:"productCodes"`
}
type FreeProductsGroups []FreeProductsGroup

func (pg *FreeProductsGroup) GetActons(order transport.IOrderRequest, tprod TableProduct) (offerentity.Actions, error) {
	result := offerentity.Actions{}
	for _, code := range pg.ProductCodes {
		product, err := tprod.GetProductByCode(code)
		if err != nil {
			continue
		}
		action := offerentity.Action{
			Type: offerentity.TypePresent,
			Data: product,
		}
		result = append(result, action)

	}
	return result, nil
}
func (pgs *FreeProductsGroups) GetActons(order transport.IOrderRequest, tprod TableProduct) (offerentity.Actions, error) {
	result := offerentity.Actions{}
	for _, pg := range *pgs {
		actions, err := pg.GetActons(order, tprod)
		if err != nil {
			continue
		}
		result = append(result, actions...)
	}

	return result, nil
}
