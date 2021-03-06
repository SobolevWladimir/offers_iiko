package offerentity

import "gopkg.in/guregu/null.v3"

type ActionCoupon struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Status   bool        `json:"status"`
	Type     int         `json:"type"`
	Comment  null.String `json:"comment"`
	Category string      `json:"category"`
	Sort     int         `json:"sort"`
}
