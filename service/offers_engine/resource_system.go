package offers_engine

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var SystemResource Resource = Resource{
	Name:  "system",
	Label: "Система",
	Attributes: Attributes{
		Attribute{
			Label: "Дата",
			Name:  "date",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "date",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						fv := fvalue.(time.Time)
						val, err := time.Parse("02-01-2006", value.(string))
						if err != nil {
							return false, err, nil
						}

						return fv.Year() == val.Year() &&
							fv.Day() == val.Day() &&
							fv.Month() == val.Month(), nil, nil

					},
				},
				Operator{
					Label: ">",
					Name:  ">",
					Input: "date",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						fv := fvalue.(time.Time)
						val, err := time.Parse("02-01-2006", value.(string))
						if err != nil {
							return false, err, nil
						}
						dat := time.Date(fv.Year(), fv.Month(), fv.Day(), 0, 0, 0, 0, time.UTC)

						return dat.After(val), nil, nil
					},
				},
				Operator{
					Label: "<",
					Name:  "<",
					Input: "date",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						fv := fvalue.(time.Time)
						val, err := time.Parse("02-01-2006", value.(string))
						if err != nil {
							return false, err, nil
						}
						dat := time.Date(fv.Year(), fv.Month(), fv.Day(), 0, 0, 0, 0, time.UTC)

						return dat.Before(val), nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Время",
			Name:  "time",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "time",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						fv := fvalue.(time.Time)
						val, err := time.Parse("15:04", value.(string))
						if err != nil {
							return false, err, nil
						}

						return fv.Hour() == val.Hour() &&
							fv.Minute() == val.Minute(), nil, nil
					},
				},
				Operator{
					Label: ">",
					Name:  ">",
					Input: "time",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						fv := fvalue.(time.Time)
						val, err := time.Parse("15:04", value.(string))
						if err != nil {
							return false, err, nil
						}
						dat := time.Date(val.Year(), val.Month(), val.Day(), fv.Hour(), fv.Minute(), val.Second(), val.Nanosecond(), val.Location())
						return dat.After(val), nil, nil
					},
				},
				Operator{
					Label: "<",
					Name:  "<",
					Input: "time",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						fv := fvalue.(time.Time)
						val, err := time.Parse("15:04", value.(string))
						if err != nil {
							return false, err, nil
						}
						dat := time.Date(val.Year(), val.Month(), val.Day(), fv.Hour(), fv.Minute(), val.Second(), val.Nanosecond(), val.Location())
						return dat.Before(val), nil, nil
					},
				},
			},
			Filters: AttributeFilters{},
		},
		Attribute{
			Label: "Запрос по url",
			Name:  "request_by_url",
			Operators: Operators{
				Operator{
					Label: "=",
					Name:  "=",
					Input: "text",
					Condition: func(fvalue, value interface{}) (bool, error, interface{}) {
						data, err := json.Marshal(fvalue)
						if err != nil {
							return false, err, nil
						}
						r := bytes.NewReader(data)
						rest, err := http.Post(value.(string), "application/json", r)
						if err != nil {
							return false, err, nil
						}
						switch rest.StatusCode {
						case 700:
							return false, nil, nil
						case 701:
							return true, nil, nil
						default:
							return false, errors.New("server returned not valid status code (need 700- false, 701- true) "), nil
						}
					},
				},
			},
			Filters: AttributeFilters{},
		},
	},
}
