package offerentity

type Allegiance struct {
	Item   RItems
	Events Events
}
type RItem struct {
	Offers Offers `json:"offers"`
	Status string `json:"status"`
}
type RItems []RItem

func (rs RItems) SetItem(item RItem) {
	for index, it := range rs {
		if it.Status == item.Status {
			rs[index] = item
			return
		}
	}
	rs = append(rs, item)
}

func (rs *RItems) SetOffers(status string, off Offers) {
	for index, it := range *rs {
		if it.Status == status {
			it.Offers = off
			(*rs)[index] = it
			return
		}
	}
	*rs = append(*rs, RItem{
		Status: status,
		Offers: off,
	})
}
func (rs *RItems) IndexOfStatus(stat string) int {
	for index, r := range *rs {
		if r.Status == stat {
			return index
		}
	}
	return -1
}
func (rs *RItems) AppendRItems(items *RItems) RItems {
	result := *rs
	for _, item := range *items {
		pos := rs.IndexOfStatus(item.Status)
		if pos == -1 {
			result = append(result, item)
		} else {
			r := result[pos]
			r.Offers.AppendOffers(&item.Offers)
			result[pos] = r
		}
	}
	return result
}
func (rs *RItems) ReplaceRItems(items *RItems) RItems {
	result := *rs
	for _, item := range *items {
		pos := result.IndexOfStatus(item.Status)
		if pos != -1 {
			result[pos] = item
		} else {
			result = append(result, item)
		}
	}
	return result
}
func (rs *RItems) GetWithoutStatuses(codes []string) RItems {
	result := RItems{}
	for _, r := range *rs {
		if !containsInArray(r.Status, codes) {
			result = append(result, r)
		}
	}
	return result
}
func (r *RItems) CreateEvent(order OfferActionInterface, stat string) Events {
	pos := r.IndexOfStatus(stat)
	if pos == -1 {
		return Events{}
	}
	item := (*r)[pos]
	result := item.Offers.CreateEvent(order, r)
	(*r)[pos] = item
	return result
}
func (r *RItems) GetCancelEvent(status string, events Events) Events {
	pos := r.IndexOfStatus(status)
	if pos == -1 {
		return Events{}
	}
	item := (*r)[pos]
	return events.GetCancelEvent(item.Offers.GetOffersIds())
}
func (r *RItems) GetCancelEventByStatuses(statuses []string, events Events) Events {
	result := Events{}
	for _, stat := range statuses {
		result.AppendEvents(r.GetCancelEvent(stat, events))

	}
	return result
}
func (r *RItems) GetActionsByType(typ int) Actions {
	result := Actions{}
	for _, item := range *r {
		result = append(result, item.GetActionsByType(typ)...)
	}

	return result
}
func (r *RItems) GetActionsSpecialDiscount() ActionsSpecialDiscount {
	result := ActionsSpecialDiscount{}
	for _, item := range *r {
		result = append(result, item.Offers.GetActionSpecialDiscount()...)
	}
	return result
}
func (r *RItems) GetActionsDiscount() ActionsDiscount {
	result := ActionsDiscount{}
	for _, item := range *r {
		result = append(result, item.Offers.GetActionsDiscount()...)
	}
	return result
}
func (r *RItem) GetActionsByType(typ int) Actions {
	return r.Offers.GetActionsByType(typ)
}
