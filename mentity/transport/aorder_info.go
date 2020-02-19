package transport

import (
	"offers_iiko/lib/base"
)

type AOrderInfo struct {
	Name          string         `json:"name"`
	Phone         string         `json:"phone"`
	OrderType     string         `json:"orderType"`
	Email         string         `json:"email"`
	TakeAwayPoint base.StringInt `json:"takeAwayPoint"`
	Comment       string         `json:"comment"`
	OrderTime     string         `json:"orderTime"`
	Time          string         `json:"time"`
	Person        int            `json:"person"`
	PayMethod     int            `json:"payMethod"`
	Cash          float32        `json:"cash"`
	NoChange      bool           `json:"noChange"`
	Promocode     string         `json:"promocode"`
	BonusPay      float32        `json:"bonusPay"`
}
