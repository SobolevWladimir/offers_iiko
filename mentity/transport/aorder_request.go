package transport

type AOrderRequest struct {
	Order     AOrder     `json:"order"`
	CityId    int        `json:"cityId"`
	OrderInfo AOrderInfo `json:"orderInfo"`
	Address   AAddress   `json:"address"`
	Platform  string     `json:"platform"`
	Token     string     `json:"token"`
}
