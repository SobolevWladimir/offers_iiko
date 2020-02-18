package offers_engine

import (
	"strconv"
	"time"
)

var ClientResource Resource = Resource{
	Label: "Клиент",
	Name:  "client",
	Attributes: Attributes{
		Attribute{
			Label: "Телефон",
			Name:  "phone",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "text",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						return fvalue == value, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Категория",
			Name:  "category",
			Operators: Operators{
				Operator{
					Label: "содержит",
					Name:  "contains",
					Input: "select",
					TypeV: "client_category",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						for _, cat := range fvalue.([]string) {
							if cat == value.(string) {
								return true, nil, nil
							}
						}
						return false, nil, nil
					},
				},
				Operator{
					Label: "не содержит",
					Name:  "not_contains",
					Input: "select",
					TypeV: "client_category",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						for _, cat := range fvalue.([]string) {
							if cat == value.(string) {
								return false, nil, nil
							}
						}
						return true, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label:   "Бонусы",
			Name:    "bonuses",
			Filters: AttributeFilters{},
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "float",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						//@fixme  когда саня сделает   поле добавить реализацию
						return false, nil, nil
					},
				},
				Operator{
					Label: ">",
					Name:  ">",
					Input: "float",
				},
				Operator{
					Label: "<",
					Name:  "<",
					Input: "float",
				},
			},
		},
		Attribute{
			Label:      "Количество заказанных товаров",
			Name:       "orders_product",
			ServerOnly: true,
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						v, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						return fvalue.(int) == v, nil, nil
					},
				},
				Operator{
					Label: "больше чем",
					Name:  ">",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						v, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						return fvalue.(int) > v, nil, nil
					},
				},
				Operator{
					Label: "меньше чем",
					Name:  "<",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						v, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						return fvalue.(int) < v, nil, nil
					},
				},
			},
			Filters: AttributeFilters{
				AttributeFilter{
					Label: "Дата  создания",
					Name:  "date_create",
					Operators: Operators{
						Operator{
							Label: "=",
							Name:  "=",
							Input: "date",
						},
						Operator{
							Label: "больше чем",
							Name:  ">",
							Input: "date",
						},
						Operator{
							Label: "меньше чем",
							Name:  "<",
							Input: "date",
						},
					},
				},
			},
		},
		Attribute{
			Label:      "Последняя активность",
			Name:       "last_activity",
			ServerOnly: true,
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "date",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := time.Parse("02-01-2006", value.(string))
						if err != nil {
							return false, err, nil
						}
						fv := fvalue.(time.Time)
						var result bool
						result = fv.Year() == val.Year() &&
							fv.Month() == val.Month() &&
							fv.Day() == val.Day()
						return result, nil, nil
					},
				},
				Operator{
					Label: "прошло дней",
					Name:  "days_passed",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						fv := fvalue.(time.Time)
						duration := time.Now().Sub(fv)
						var result bool
						result = duration.Round(time.Hour*24).Hours() == float64(val*24)
						return result, nil, nil
					},
				},
				Operator{
					Label: "прошло дней больше чем",
					Name:  "days_passed_more",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						fv := fvalue.(time.Time)
						duration := time.Now().Sub(fv)
						var result bool
						result = duration.Round(time.Hour*24).Hours() > float64(val*24)
						return result, nil, nil
					},
				},
				Operator{
					Label: "прошло дней меньше чем",
					Name:  "days_passed_less",
					Input: "int",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						val, err := strconv.Atoi(value.(string))
						if err != nil {
							return false, err, nil
						}
						fv := fvalue.(time.Time)
						duration := time.Now().Sub(fv)
						var result bool
						result = duration.Round(time.Hour*24).Hours() < float64(val*24)
						return result, nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label:      "Количество заказов",
			Name:       "order_count",
			ServerOnly: true,
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
					Label: ">",
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
					Label: "<",
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
			Filters: AttributeFilters{
				AttributeFilter{
					Label: "Дата",
					Name:  "date_create",
					Operators: Operators{
						Operator{
							Label: "=",
							Name:  "=",
							Input: "date",
						},
						Operator{
							Label: "больше чем",
							Name:  ">",
							Input: "date",
						},
						Operator{
							Label: "меньше чем",
							Name:  "<",
							Input: "date",
						},
					},
				},
				AttributeFilter{
					Label: "Промокод",
					Name:  "coupon",
					Operators: Operators{
						Operator{
							Label: "=",
							Name:  "=",
							Input: "text",
						},
					},
				},
				AttributeFilter{
					Label: "Категория Промокода",
					Name:  "coupon_category",
					Operators: Operators{
						Operator{
							Label: "входит",
							Name:  "contains",
							Input: "select",
							TypeV: "coupon_category",
						},
					},
				},
			},
		},
	},
}
