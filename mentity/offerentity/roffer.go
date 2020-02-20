package offerentity

type Offer struct {
	Id      string  `json:"id"`
	Error   bool    `json:"error"`
	Message string  `json:"message"`
	Actions Actions `json:"actions"`
}
type Offers []Offer

func (of *Offers) IndexOfId(id string) int {
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
