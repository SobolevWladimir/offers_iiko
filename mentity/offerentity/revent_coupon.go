package offerentity

//Списывались ли купоны
type EventCoupon struct {
	Completed bool   `json:"completed"` //выполнено ли действие
	Coupon    string `json:"coupons"`   //Идентификтор купона
	Status    bool   `json:"status"`    // какой статус установить
	Note      string `json:"note"`
	OfferId   int    `json:"offer_id"` //идентификатор акции на котором сработал
}
type EventsCoupon []EventCoupon

func (bs *EventsCoupon) SetOfferId(id int) {
	for index, b := range *bs {
		b.OfferId = id
		(*bs)[index] = b
	}
}
func (ev *EventsCoupon) GetCancelEvent(offers []int) EventsCoupon {
	result := EventsCoupon{}
	for _, of := range offers {
		for i := len(*ev) - 1; i >= 0; i-- {
			item := (*ev)[i]
			if item.OfferId == of {
				result = append(result, EventCoupon{
					Coupon:  item.Coupon,
					Status:  true,
					Note:    " cancel",
					OfferId: item.OfferId,
				})
				break
			}
		}
	}
	return result
}
func (bs *EventsCoupon) SetCompleted(flag bool) {
	for index, b := range *bs {
		b.Completed = flag
		(*bs)[index] = b
	}
}

func (ev *EventsCoupon) FilterByCompleted(flag bool) EventsCoupon {
	result := EventsCoupon{}
	for _, e := range *ev {
		if e.Completed == flag {
			result = append(result, e)
		}
	}
	return result
}
func (ev *EventsCoupon) FilterByStatus(flag bool) EventsCoupon {
	result := EventsCoupon{}
	for _, e := range *ev {
		if e.Status == flag {
			result = append(result, e)
		}
	}
	return result
}
func (ev *EventsCoupon) GetCouponIds() []string {
	result := []string{}
	for _, e := range *ev {
		result = append(result, e.Coupon)
	}
	return result
}
