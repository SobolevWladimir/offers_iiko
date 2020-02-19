package transport

import "strconv"

type AAddress struct {
	Street   AStreet `json:"street"`
	Building string  `json:"building"`
	Entrance int     `json:"entrance"`
	Floor    int     `json:"floor"`
	Room     int     `json:"room"`
}

func (a *AAddress) GetIAddress() IAddress {
	return IAddress{
		Street:    a.Street.Name,
		Home:      a.Building,
		Entrance:  strconv.Itoa(a.Entrance),
		Floor:     strconv.Itoa(a.Floor),
		Apartment: strconv.Itoa(a.Room),
	}

}
