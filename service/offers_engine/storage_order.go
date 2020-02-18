package offers_engine

const InputNameOrderDiliveryType = "dilivery_type"
const InputNameOrdetPoints = "points"

func getOrdetDelivetyType() InputValues {
	result := InputValues{
		InputValue{
			Label: "Самовывоз",
			Value: "0",
		},
		InputValue{
			Label: "Доставка",
			Value: "1",
		},
		InputValue{
			Label: "В зале",
			Value: "2",
		},
	}
	return result
}
