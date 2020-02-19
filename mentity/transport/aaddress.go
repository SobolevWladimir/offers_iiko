package transport

type StreetEntity struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	TypeShort string `json:"typeShort"`
	Code      string `json:"code"`
	Status    int    `json:"status"`
	Sity      int    `json:"sity"`
}
type AAddress struct {
	Street   StreetEntity `json:"street"`
	Building string       `json:"building"`
	Entrance int          `json:"entrance"`
	Floor    int          `json:"floor"`
	Room     int          `json:"room"`
}
