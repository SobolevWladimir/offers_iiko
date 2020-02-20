package iiko

import (
	"offers_iiko/mentity/offerentity"
	"offers_iiko/mentity/transport"
)

type CheckinResult struct {
	LoyatyResult       LoyaltyResult    `json:"loyatyResult"`
	ValidationWarnings ValidatorWarning `json:"validationWarnings"`
}

func (c *CheckinResult) GetActons(order transport.IOrderRequest) (offerentity.Actions, error) {
	return c.LoyatyResult.ProgramResults.GetActons(order)
}
