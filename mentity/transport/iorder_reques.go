package transport

type IOrderRequest struct {
	Organization string    `json:"organization"`
	Customer     ICustomer `json:"customer"`
	Order        IOrder    `json:"order"`
	Coupon       string    `json:"coupon"`
}
