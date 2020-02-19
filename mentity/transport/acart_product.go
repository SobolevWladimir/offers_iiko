package transport

type ACartProducts []ACartProduct
type ACartProduct struct {
	Product  ProductItem  `json:"product"`
	Quantity int          `json:"quantity"`
	Options  ACartOptions `json:"options"`
}