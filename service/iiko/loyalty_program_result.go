package iiko

import (
	"offers_iiko/mentity/offerentity"
	"offers_iiko/mentity/transport"
)

type LoyaltyProgramResult struct {
	ProgramId       string             `json:"programId"`
	Name            string             `json:"name"`
	Discounts       DiscountOperations `json:"discounts"`       //Скидки, примененные к отдельным элементам заказа
	Upsales         Upsales            `json:"upsales"`         //Список доступных для заказа подсказок
	FreeProducts    FreeProductsGroups `json:"freeProducts"`    // Список доступных для заказа подарочных блюд
	Combos          []string           `json:"combos"`          // Список доступных для добавления комбо
	AvailableCombos AvailableCombos    `json:"availableCombos"` // Список доступных для досборки комбо

}
type LoyaltyProgramResults []LoyaltyProgramResult

func (l *LoyaltyProgramResults) GetActons(order transport.IOrderRequest, tprod TableProduct) (offerentity.Actions, error) {
	result := offerentity.Actions{}
	for _, lo := range *l {
		actions, err := lo.GetActons(order, tprod)
		if err != nil {
			return result, err
		}
		result = append(result, actions...)
	}

	return result, nil
}
func (l *LoyaltyProgramResult) GetActons(order transport.IOrderRequest, tprod TableProduct) (offerentity.Actions, error) {
	result := offerentity.Actions{}
	// discount
	dis, err := l.Discounts.GetActons(order)
	if err != nil {
		return result, err
	}
	result = append(result, dis...)
	//upsale
	result = append(result, l.Upsales.GetActons()...)
	//free products
	free, err := l.FreeProducts.GetActons(order, tprod)

	result = append(result, free...)

	return result, err
}
