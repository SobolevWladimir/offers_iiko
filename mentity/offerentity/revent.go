package offerentity

type Events struct {
	Bonuses        EventsBonuses        `json:"bonuses"`
	Coupons        EventsCoupon         `json:"coupons"`
	ClientCategory EventsClientCategory `json:"client_category"`
}

func (ev *Events) AppendEvents(val Events) {
	ev.Bonuses = append(ev.Bonuses, val.Bonuses...)
	ev.Coupons = append(ev.Coupons, val.Coupons...)
	ev.ClientCategory = append(ev.ClientCategory, val.ClientCategory...)
}
func (ev *Events) SetOfferId(id int) {
	ev.Bonuses.SetOfferId(id)
	ev.ClientCategory.SetOfferId(id)
	ev.Coupons.SetOfferId(id)
}
func (ev *Events) GetCancelEvent(offers []int) Events {
	result := Events{}
	result.Bonuses = ev.Bonuses.GetCancelEvent(offers)
	result.ClientCategory = ev.ClientCategory.GetCancelEvent(offers)
	result.Coupons = ev.Coupons.GetCancelEvent(offers)
	return result
}
