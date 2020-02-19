package transport

import (
	"offers_iiko/lib/base"
	"time"
)

type AOrderRequest struct {
	Order     AOrder     `json:"order"`
	CityId    int        `json:"cityId"`
	OrderInfo AOrderInfo `json:"orderInfo"`
	Address   AAddress   `json:"address"`
	Platform  string     `json:"platform"`
	Token     string     `json:"token"`
}

func (o *AOrderRequest) ToIOrderRequest() (IOrderRequest, error) {
	result := IOrderRequest{}
	result.Customer = o.OrderInfo.GetICustomer()
	result.Coupon = o.OrderInfo.Promocode
	order, err := o.GetIOrder()

	result.Order = order
	return result, err
}
func (o *AOrderRequest) GetIOrder() (IOrder, error) {
	result := IOrder{}
	result.Date = IDateTimeUTC(time.Now())
	result.ID = base.UUID()
	result.Phone = o.OrderInfo.GetClearPhone()
	result.IsSelfService = o.OrderInfo.OrderType == "takeAway"
	result.Address = o.Address.GetIAddress()
	result.Comment = o.OrderInfo.Comment
	result.PersonCount = o.OrderInfo.Person
	result.FullSumm = o.Order.TotalPrice
	result.MarketingSource = o.Platform + " check"

	return result, nil
}
