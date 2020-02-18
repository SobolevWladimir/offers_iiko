package offerentity

// Измененые  категории у клиента
type EventClientCategory struct {
	Completed       bool     `json:"completed"` //выполнено ли действие
	CategoryRemoved []string `json:"category_removed"`
	CategoryAdded   []string `json:"category_added"`
	Note            string   `json:"note"`
	OfferId         int      `json:"offer_id"` //идентификатор акции на котором сработал
}
type EventsClientCategory []EventClientCategory

func (bs *EventsClientCategory) SetOfferId(id int) {
	for index, b := range *bs {
		b.OfferId = id
		(*bs)[index] = b
	}
}
func (ev *EventsClientCategory) GetCancelEvent(offers []int) EventsClientCategory {
	result := EventsClientCategory{}
	for _, of := range offers {
		for i := len(*ev) - 1; i >= 0; i-- {
			item := (*ev)[i]
			if item.OfferId == of {
				catadded := item.CategoryRemoved
				catdel := item.CategoryAdded
				result = append(result, EventClientCategory{
					CategoryAdded:   catadded,
					CategoryRemoved: catdel,
					Note:            " cancel",
					OfferId:         item.OfferId,
				})
				break
			}
		}

	}
	return result
}

func (bs *EventsClientCategory) SetCompleted(flag bool) {
	for index, b := range *bs {
		b.Completed = flag
		(*bs)[index] = b
	}
}

func (ev *EventsClientCategory) FilterByCompleted(flag bool) EventsClientCategory {
	result := EventsClientCategory{}
	for _, e := range *ev {
		if e.Completed == flag {
			result = append(result, e)
		}
	}
	return result
}
func (ev *EventsClientCategory) GetAddedCategory() []string {
	result := []string{}
	for _, e := range *ev {
		result = append(result, e.CategoryAdded...)
	}
	return result
}
func (ev *EventsClientCategory) GetRemovedCategory() []string {
	result := []string{}
	for _, e := range *ev {
		result = append(result, e.CategoryRemoved...)
	}
	return result

}
