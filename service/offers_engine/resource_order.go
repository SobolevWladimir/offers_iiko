package offers_engine

import (
	"altegra_offers/service/coupon"
	"strconv"
)

var OrderResource Resource = Resource{
	Name:  "order",
	Label: "Заказ",
	Attributes: Attributes{
		Attribute{
			Label: "Тип",
			Name:  "delivery",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "select",
					TypeV: InputNameOrderDiliveryType,
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						return fvalue.(int) == val, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Точка",
			Name:  "point",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "select",
					TypeV: InputNameOrdetPoints,
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						return fvalue == value, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Купон",
			Name:  "coupon",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "text",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						if fvalue == nil {
							return false, nil, nil
						}
						entity := fvalue.(coupon.Coupon)
						return entity.Status && entity.Name == value.(string), nil, entity
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Количество персон",
			Name:  "person",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						return fvalue.(int) == val, nil, nil
					},
				},
				Operator{
					Label: "больше чем",
					Name:  ">",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						return fvalue.(int) > val, nil, nil
					},
				},
				Operator{
					Label: "меньше чем",
					Name:  "<",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						return fvalue.(int) < val, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Сумма заказа без учета скидок",
			Name:  "pre_amount",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "float",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.ParseFloat(value.(string), 32)
						fv := fvalue.(float32)

						if err != nil {
							return false, err, nil
						}
						return fv == float32(val), nil, nil
					},
				},
				Operator{
					Label: ">",
					Name:  ">",
					Input: "float",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.ParseFloat(value.(string), 32)
						fv := fvalue.(float32)

						if err != nil {
							return false, err, nil
						}
						return fv > float32(val), nil, nil
					},
				},
				Operator{
					Label: "<",
					Name:  "<",
					Input: "float",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.ParseFloat(value.(string), 32)
						fv := fvalue.(float32)

						if err != nil {
							return false, err, nil
						}
						return fv < float32(val), nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Сумма заказа",
			Name:  "amount",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "float",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.ParseFloat(value.(string), 32)
						fv := fvalue.(float32)

						if err != nil {
							return false, err, nil
						}
						return fv == float32(val), nil, nil
					},
				},
				Operator{
					Label: ">",
					Name:  ">",
					Input: "float",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.ParseFloat(value.(string), 32)
						fv := fvalue.(float32)

						if err != nil {
							return false, err, nil
						}
						return fv > float32(val), nil, nil
					},
				},
				Operator{
					Label: "<",
					Name:  "<",
					Input: "float",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.ParseFloat(value.(string), 32)
						fv := fvalue.(float32)

						if err != nil {
							return false, err, nil
						}
						return fv < float32(val), nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Маркеры",
			Name:  "markers",
			Operators: Operators{
				Operator{
					Label: "содержат",
					Name:  "contains",
					Input: "select",
					TypeV: "order_markers",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						ms := fvalue.([]string)
						for _, m := range ms {
							if m == value {
								return true, nil, nil
							}
						}
						return false, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Продукты",
			Name:  "products",
			Operators: Operators{
				Operator{
					Label: "содержат",
					Name:  "contains",
					Input: "select",
					TypeV: "products",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						pr := fvalue.([]string)
						for _, p := range pr {
							if p == value {
								return true, nil, nil
							}
						}
						return false, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
	},
}
