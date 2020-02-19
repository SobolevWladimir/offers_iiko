package transport

type IOrderItem struct {
	ID          string              `json:"id"`
	OrderItemId string              `json:"OrderItemId"`
	Cagegory    string              `json:"cagegory"`
	Code        string              `json:"code"`
	Name        string              `json:"name"`
	Amount      string              `json:"amount"`
	Sum         float32             `json:"sum"`
	Modifiers   IOrderItemModifiers `json:"modifiers"`
	GuesId      string              `json:"GuesId"`
}
