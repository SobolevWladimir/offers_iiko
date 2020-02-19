package transport

import "time"

type ATime time.Time
type AOrderInfo struct {
	Name          string  `json:"name"`
	Phone         string  `json:"phone"`
	OrderType     string  `json:"orderType"`
	Email         string  `json:"email"`
	TakeAwayPoint string  `json:"takeAwayPoint"`
	Comment       string  `json:"comment"`
	OrderTime     string  `json:"orderTime"`
	Time          ATime   `json:"time"`
	Person        int     `json:"person"`
	PayMethod     int     `json:"payMethod"`
	Cash          float32 `json:"cash"`
	NoChange      bool    `json:"noChange"`
	Promocode     string  `json:"promocode"`
	BonusPay      string  `json:"bonusPay"`
}
