package transport

type IAddress struct {
	City               string `json:"city"`
	Street             string `json:"street"`
	StreetId           string `json:"StreetId"`
	StreetClassifierId string `json:"StreetClassifierId"`
	Home               string `json:"home"`
	Housing            string `json:"housing"`
	Apartment          string `json:"apartment"`
	Entrance           string `json:"entrance"`
	Floor              string `json:"floor"`
	Doorphone          string `json:"doorphone"`
	Comment            string `json:"comment"`
}
