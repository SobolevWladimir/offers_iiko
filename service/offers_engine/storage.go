package offers_engine

import (
	"altegra_offers/service/client_category"
	"altegra_offers/service/coupon_category"
	"strconv"
)

type InputValue struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
type InputValues []InputValue
type Input struct {
	Name   string      `json:"name"`
	Values InputValues `json:"values"`
}
type Inputs []Input
type Storage struct {
	Resources   Resources `json:"resource"`
	InputValues Inputs    `json:"input_values"`
}

func GetStandartStorage(city int) Storage {
	inputs := Inputs{
		Input{
			Name:   "coupon_category",
			Values: getCouponCategorys(city),
		},
		Input{
			Name:   InputNameOrderDiliveryType,
			Values: getOrdetDelivetyType(),
		},
	}
	return Storage{
		Resources:   StandartResources,
		InputValues: inputs,
	}
}

var StandartResources Resources = Resources{
	&SystemResource,
	&ClientResource,
	&OrderResource,
}

func getClientCategorys() InputValues {
	result := InputValues{}
	cats, err := client_category.FindAll()
	if err != nil {
		return result
	}
	for _, cat := range cats {
		result = append(result, InputValue{
			Label: cat.Name,
			Value: cat.Name,
		})
	}
	return result
}
func getCouponCategorys(city int) InputValues {
	result := InputValues{}

	cats, err := coupon_category.FindByCity(city)
	if err != nil {
		return result
	}
	for _, cat := range cats {
		val := strconv.Itoa(cat.Id)
		result = append(result, InputValue{
			Label: cat.Name,
			Value: val,
		})
	}
	return result
}
