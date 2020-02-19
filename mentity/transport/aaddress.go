package transport

type AAddress struct {
	Street   AStreet `json:"street"`
	Building string  `json:"building"`
	Entrance int     `json:"entrance"`
	Floor    int     `json:"floor"`
	Room     int     `json:"room"`
}
