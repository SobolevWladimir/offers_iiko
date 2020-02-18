package offerentity

type Offer struct {
	Id      int     `json:"id"`
	Error   bool    `json:"error"`
	Message string  `json:"message"`
	Actions Actions `json:"actions"`
}
type Offers []Offer

func (of *Offers) IndexOfId(id int) int {
	for index, o := range *of {
		if o.Id == id {
			return index
		}
	}
	return -1
}
func (of *Offers) AppendOffers(val *Offers) {
	for index, o := range *of {
		pos := val.IndexOfId(o.Id)
		//	если акция  уже есть в массиве
		if pos != -1 {
			ff := (*val)[pos]
			if ff.Error != true {
				(*of)[index] = ff
			}
		} else {
			*of = append(*of, *val...)
		}
	}
}
func (of *Offers) CreateEvent(order OfferActionInterface, r *RItems) Events {
	result := Events{}
	for index, o := range *of {
		result.AppendEvents(o.CreateEvent(order, r))
		(*of)[index] = o
	}
	return result
}
func (o *Offer) CreateEvent(order OfferActionInterface, r *RItems) Events {
	events := o.Actions.CreateEvent(order, r)
	events.SetOfferId(o.Id)
	return events
}
func (os *Offers) GetOffersIds() []int {
	result := []int{}
	for _, o := range *os {
		result = append(result, o.Id)
	}
	return result
}

func (of *Offers) GetActionsByType(typ int) Actions {
	result := Actions{}
	for _, o := range *of {
		result = append(result, o.Actions.GetActionsByType(typ)...)
	}
	return result
}
func (of *Offers) GetActionSpecialDiscount() ActionsSpecialDiscount {
	result := ActionsSpecialDiscount{}
	for _, o := range *of {
		result = append(result, o.Actions.GetSpecialDiscount()...)
	}
	return result
}
func (of *Offers) GetActionsDiscount() ActionsDiscount {
	result := ActionsDiscount{}
	for _, o := range *of {
		result = append(result, o.Actions.GetDiscount()...)
	}
	return result
}
