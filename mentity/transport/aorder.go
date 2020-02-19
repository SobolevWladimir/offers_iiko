package transport

type AOrder struct {
	TotalPrice float32       `json:"totalPrice"`
	Products   ACartProducts `json:"products"`
}
