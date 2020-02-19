package transport

import "fmt"

type StreetEntity struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	TypeShort string `json:"typeShort"`
	Code      string `json:"code"`
	Status    int    `json:"status"`
	Sity      int    `json:"sity"`
}
type AStreet StreetEntity

func (a *AStreet) UnmarshalJSON(data []byte) error {
	fmt.Println("street parse :", string(data))
	return nil
}
