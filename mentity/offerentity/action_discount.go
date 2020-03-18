package offerentity

type ActionDiscount struct {
	SaleType int                      `json:"sale_type"`
	Type     TypeValueSpecialDiscount `json:"type"`
	Value    float32                  `json:"value"`
}
type ActionsDiscount []ActionDiscount
