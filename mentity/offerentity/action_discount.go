package offerentity

type ActionDiscount struct {
	SaleType int                      `json:"sale_type"`
	Type     TypeValueSpecialDiscount `json:"type"`
	Value    float32                  `josn:"value"`
}
type ActionsDiscount []ActionDiscount
