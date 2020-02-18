package offerentity

type ActionClientCategory struct {
	Category string `json:"category"`
	Type     int    `json:"type"`
}

func (a *ActionClientCategory) CreateEvent(order OfferActionInterface, r *RItems) EventClientCategory {
	result := EventClientCategory{}
	switch a.Type {
	case 0:
		result.CategoryAdded = append(result.CategoryAdded, a.Category)
	case 1:
		result.CategoryAdded = append(result.CategoryRemoved, a.Category)
	}
	return result

}
