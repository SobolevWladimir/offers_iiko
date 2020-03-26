package transport

import (
	"offers_iiko/lib/base"
	"time"
)

type AOrderRequest struct {
	Order     AOrder     `json:"order"`
	CityId    int        `json:"cityId"`
	OrderInfo AOrderInfo `json:"orderInfo"`
	Address   AAddress   `json:"-"`
	Platform  string     `json:"platform"`
	Token     string     `json:"token"`
}
type TableInterface interface {
	FindCityNameById(id int) (string, error)
}
type NomenclatureInterface interface {
	FindCategoryNameByProductCode(code string) (string, error)
	FindProductIdByCode(code string) (string, error)
}

func (o *AOrderRequest) ToIOrderRequest(db TableInterface, nom NomenclatureInterface) (IOrderRequest, error) {
	result := IOrderRequest{}
	result.Customer = o.OrderInfo.GetICustomer()
	result.Coupon = o.OrderInfo.Promocode
	order, err := o.GetIOrder(db, nom)
	result.Order = order
	return result, err
}
func (o *AOrderRequest) GetIOrder(db TableInterface, nom NomenclatureInterface) (IOrder, error) {
	result := IOrder{}
	result.Date = IDateTimeUTC(time.Now())
	result.ID = base.UUID()
	result.Phone = o.OrderInfo.GetClearPhone()
	result.IsSelfService = o.OrderInfo.OrderType != "delivery"
	city, err := db.FindCityNameById(o.CityId)
	if err != nil {
		return result, err
	}
	result.Address = o.Address.GetIAddress(city)
	result.Comment = o.OrderInfo.Comment
	result.PersonCount = o.OrderInfo.Person
	result.FullSumm = o.Order.TotalPrice
	result.MarketingSource = o.Platform + " check"
	result.Items = o.Order.Products.ToOrderItems(nom)
	return result, nil
}
