package offerentity

type Action struct {
	IsDone int        `json:"is_done"` // выполнено ли дейсвие 0-не выполнена 1 - выполнена на сервере 2-выполнена клиентом
	Type   int        `json:"type"`    //	 тип действия
	Data   ActionData `json:"data"`    // Данные действия
}
type Actions []Action

//типы действий
const (
	TypePresent              = 1
	TypeSpecialDiscount      = 2
	TypeDiscount             = 3
	TypeActionBonuses        = 7
	TypeActionClientCategory = 8
	TypeActionCoupon         = 1000
)

func (a *Action) CreateEvent(order OfferActionInterface, r *RItems) Events {
	result := Events{}
	switch a.Type {
	case TypeActionBonuses:
		{
			entity := a.Data.ToActinBonuses()
			event := entity.CreateEvent(order, r)
			result.Bonuses = append(result.Bonuses, event)
			a.IsDone = 1
		}
	case TypeActionClientCategory:
		{
			entity := a.Data.ToActionClientCategory()
			event := entity.CreateEvent(order, r)
			result.ClientCategory = append(result.ClientCategory, event)
			a.IsDone = 1
		}
	case TypeActionCoupon:
		{
			entity := a.Data.ToActinCoupon()
			event := entity.CreateEvent(order, r)
			result.Coupons = append(result.Coupons, event)
			a.IsDone = 1
		}
	}
	return result
}
func (as *Actions) CreateEvent(order OfferActionInterface, r *RItems) Events {
	result := Events{}
	for index, a := range *as {
		result.AppendEvents(a.CreateEvent(order, r))
		(*as)[index] = a
	}
	return result
}
func (ofs *Actions) GetIsDone() Actions {
	result := Actions{}
	for _, of := range *ofs {
		if of.IsDone == 1 {
			result = append(result, of)
		}

	}
	return result
}
func (ac *Actions) GetActionsByType(typ int) Actions {
	result := Actions{}
	for _, a := range *ac {
		if a.Type == typ {
			result = append(result, a)
		}
	}
	return result
}
func (ac *Actions) GetSpecialDiscount() ActionsSpecialDiscount {
	result := ActionsSpecialDiscount{}
	for _, a := range *ac {
		if a.Type == TypeSpecialDiscount {
			result = append(result, a.Data.ToActionSpecialDiscount())
		}
	}

	return result
}
func (ac *Actions) GetDiscount() ActionsDiscount {
	result := ActionsDiscount{}
	for _, a := range *ac {
		if a.Type == TypeDiscount {
			result = append(result, a.Data.ToActionDiscount())
		}
	}

	return result
}
