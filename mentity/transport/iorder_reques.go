package transport

type IOrderRequest struct {
	Organization string
	Customer     ICustomer
	Order        IOrder
	Coupon       string
}
