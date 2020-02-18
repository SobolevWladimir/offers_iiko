package offerentity

//Калькурировались ли баллы клиента
type EventBonuses struct {
	Completed bool    `json:"completed"` //выполнено ли действие
	Value     float32 `json:"value"`     //Сколько списанно
	Note      string  `json:"note"`
	OfferId   int     `json:"offer_id"` //идентификатор акции на котором сработал
}
type EventsBonuses []EventBonuses

func (ev *EventsBonuses) GetCancelEvent(offers []int) EventsBonuses {
	result := EventsBonuses{}
	for _, of := range offers {
		for i := len(*ev) - 1; i >= 0; i-- {
			item := (*ev)[i]
			if item.OfferId == of {
				result = append(result, EventBonuses{
					Value:   -item.Value,
					Note:    " cancel",
					OfferId: item.OfferId,
				})
				break
			}
		}
	}
	return result
}
func (bs *EventsBonuses) SetOfferId(id int) {
	for index, b := range *bs {
		b.OfferId = id
		(*bs)[index] = b
	}
}
func (bs *EventsBonuses) SetCompleted(flag bool) {
	for index, b := range *bs {
		b.Completed = flag
		(*bs)[index] = b
	}
}
func (ev *EventsBonuses) SummarizeValue() float32 {
	var result float32 = 0
	for _, e := range *ev {
		result = result + e.Value
	}
	return result
}
func (ev *EventsBonuses) FilterByCompleted(flag bool) EventsBonuses {
	result := EventsBonuses{}
	for _, e := range *ev {
		if e.Completed == flag {
			result = append(result, e)
		}
	}
	return result
}

func (ev *EventsBonuses) CalculateValue(completed bool) float32 {
	items := ev.FilterByCompleted(completed)
	return items.SummarizeValue()
}
