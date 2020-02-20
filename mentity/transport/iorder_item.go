package transport

type IOrderItem struct {
	ID          string              `json:"id"`
	OrderItemId string              `json:"-"`
	Cagegory    string              `json:"-"`
	Code        string              `json:"code"`
	Name        string              `json:"-"`
	Amount      float32             `json:"amount"`
	Sum         float32             `json:"sum"`
	Modifiers   IOrderItemModifiers `json:"-"`
	GuesId      string              `json:"-"`
}
