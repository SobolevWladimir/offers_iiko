package transport

type IOrderItem struct {
	ID          string              `json:"id"`
	OrderItemId string              `json:"orderItemId"`
	Cagegory    string              `json:"category"`
	Code        string              `json:"code"`
	Name        string              `json:"name"`
	Amount      float32             `json:"amount"`
	Sum         float32             `json:"sum"`
	SiteId      int                 `json:"-"`
	Modifiers   IOrderItemModifiers `json:"-"`
	GuesId      string              `json:"-"`
}
