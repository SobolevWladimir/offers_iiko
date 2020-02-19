package transport

type IOrderItemModifier struct {
	ID        string  `json:"id"`
	OrderItem string  `json:"orderItem"`
	Name      string  `json:"name"`
	Amount    float32 `json:"amount"`
	GroupId   string  `json:"groupId"`
	GroupName string  `json:"groupName"`
}
