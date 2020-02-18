package offerentity

type ActionBonuses struct {
	All       bool     `json:"all"`
	Exception []string `json:"exception"`
	Type      int      `json:"type"`
	Value     float32  `jso:"value"`
}

func (a *ActionBonuses) CreateEvent(order OfferActionInterface, r *RItems) EventBonuses {
	if a.Type == 0 {
		return EventBonuses{
			Completed: false,
			Value:     a.Value,
		}
	}
	if a.Type == 1 {
		// расчет в процентах
		bonus, err := order.CalculateClientBonuses(a, r)
		var note string
		if err != nil {
			note = " can't calculate  client bonus:" + err.Error()
		}
		return EventBonuses{
			Completed: false,
			Value:     bonus,
			Note:      note,
		}

	}
	return EventBonuses{
		Completed: true,
		Value:     0,
		Note:      "unknown type",
	}
}
