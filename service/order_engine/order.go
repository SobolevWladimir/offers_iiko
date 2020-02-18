package order_engine

import (
	"altegra_offers/lib/base"
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/order_engine/address"
	"altegra_offers/service/order_engine/mlink"
	"altegra_offers/service/order_engine/orders"
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v3"
)

//@fixme add offers event
type Order struct {
	Id             string             `json:"id"  valid:"uuid"`
	Point          string             `json:"point" valid:"uuid"`
	Client         null.String        `json:"client" valid:"-"`
	Parent         null.String        `json:"parent" valid:"-"`
	Status         string             `json:"status" valid:"-"`
	Type           int                `json:"type" valid:"-"`
	Delivery       int                `json:"delivery" valid:"-"`
	Person         int                `json:"person" valid:"-"`
	CookInDate     base.DateFP        `json:"cook_in_date" valid:"-"`
	CookInTime     base.NullTimeFP    `json:"cook_in_time" valid:"-"`
	Paid           bool               `json:"paid" valid:"-"`
	LocalNumber    null.String        `json:"local_number" valid:"-"`
	Comment        null.String        `json:"comment" valid:"-"`
	PreAmount      float32            `json:"pre_amount" valid:"-"`
	Amount         float32            `json:"amount" valid:"-"`
	PersonInCharge null.String        `json:"person_in_charge" valid:"-"`
	Offers         offerentity.RItems `json:"offers" valid:"-"`
	OfferEvent     offerentity.Events `json:"offers_event" valid:"-"`
	LastUpdate     base.DateTimeFP    `json:"last_update" valid:"-"`
	Created        base.DateTimeFP    `json:"created" valid:"-"`
	Address        NullAddress        `json:"address" valid:"-"`
	Coupon         null.String        `json:"coupon" valid:"-"`
	Markers        []string           `json:"markers" valid:"-"`
	Payments       Payments           `json:"payments" valid:"-"`
	Products       Products           `json:"products" valid:"-"`
}

type Orders []Order

func toOrder(or *orders.Order) (Order, error) {
	result := Order{}
	result.Id = or.Id
	result.Point = or.Point
	result.Client = or.Client
	result.Parent = or.Parent
	result.Status = or.Status
	result.Type = or.Type
	result.Delivery = or.Delivery
	result.Person = or.Person
	result.CookInDate = base.DateFP(or.CookInDate)

	result.CookInTime = base.NewNullTime(or.CookInTime)
	result.Paid = or.Paid
	result.LocalNumber = or.LocalNumber
	result.Comment = or.Comment
	result.PreAmount = or.PreAmount
	result.Amount = or.Amount
	result.PersonInCharge = or.PersonInCharge
	result.Coupon = or.Coupon
	// @fixme  convent offers
	//result.OfferEvent
	result.LastUpdate = base.DateTimeFP(or.LastUpdate)
	result.Created = base.DateTimeFP(or.Created)

	json.Unmarshal([]byte(or.Offers.String), &result.Offers)
	json.Unmarshal([]byte(or.OfferEvent.String), &result.OfferEvent)

	prods, err := FindAllProductByOrder(or.Id)
	if err != nil {
		return result, err
	}
	result.Products = prods
	if markers, err := mlink.FindAllByOrder(or.Id); err != nil {
		return result, err
	} else {
		result.Markers = markers
	}
	payments, err := FindAllPaymentsByOrders(or.Id)
	if err != nil {
		return result, err
	}
	result.Payments = payments
	if ad, err := FindAddressByOrder(or.Id); err == nil {
		result.Address.Address = ad
		result.Address.Valid = true
	}
	return result, nil
}
func (or *Order) toDBEntity() orders.Order {
	result := orders.Order{}
	result.Id = or.Id
	result.Point = or.Point
	result.Client = or.Client
	result.Parent = or.Parent
	result.Status = or.Status
	result.Type = or.Type
	result.Delivery = or.Delivery
	result.Person = or.Person
	result.CookInDate = time.Time(or.CookInDate)
	result.CookInTime = or.CookInTime.ToTime()
	result.Paid = or.Paid
	result.LocalNumber = or.LocalNumber
	result.Comment = or.Comment
	result.PreAmount = or.PreAmount
	result.Amount = or.Amount
	result.PersonInCharge = or.PersonInCharge
	result.Coupon = or.Coupon
	result.LastUpdate = time.Time(or.LastUpdate)
	result.Created = time.Time(or.Created)
	offers, err := json.Marshal(or.Offers)
	result.Offers = null.NewString(string(offers), err == nil)
	o_event, err := json.Marshal(or.OfferEvent)
	result.OfferEvent = null.NewString(string(o_event), err == nil)
	return result
}
func toOrders(ors *orders.Orders) (Orders, error) {
	result := Orders{}
	for _, or := range *ors {
		order, err := toOrder(&or)
		if err != nil {
			return result, err
		}
		result = append(result, order)
	}
	return result, nil
}
func FindOrders(date_start, date_end string) (Orders, error) {
	orders, err := orders.FindAll(date_start, date_end)
	if err != nil {
		return Orders{}, err
	}
	return toOrders(&orders)
}
func FindOrdersByPoint(date_start, date_end, point string) (Orders, error) {
	orders, err := orders.FindAllByPoint(date_start, date_end, point)
	if err != nil {
		return Orders{}, err
	}
	return toOrders(&orders)
}
func FindOrderById(id string) (Order, error) {
	order, err := orders.FindOneById(id)
	if err != nil {
		return Order{}, err
	}
	return toOrder(&order)
}
func InsertOrder(order *Order) error {
	dbentity := order.toDBEntity()
	if err := orders.Insert(&dbentity); err != nil {
		return fmt.Errorf("insert order entity:%v", err.Error())
	}
	if order.Address.Valid {
		if err := order.Address.Address.Insert(order.Id); err != nil {
			return fmt.Errorf("insert address:%v", err.Error())
		}
	}
	if err := mlink.InsertArray(order.Markers, order.Id); err != nil {
		return fmt.Errorf("insert markers:%v", err.Error())
	}
	if err := order.Payments.Insert(order.Id); err != nil {
		return fmt.Errorf("insert payments:%v", err.Error())
	}
	if err := order.Products.Insert(order.Id); err != nil {
		return fmt.Errorf("insert product:%v", err.Error())
	}
	return nil
}
func SaveOrder(order *Order) error {
	tx := dataBase.MustBegin()
	dbentity := order.toDBEntity()
	if err := orders.TxSave(tx, &dbentity); err != nil {
		tx.Rollback()
		return fmt.Errorf("save order entity:%v", err.Error())
	}
	if order.Address.Valid {
		if err := order.Address.Address.TxSave(tx, order.Id); err != nil {
			tx.Rollback()
			return fmt.Errorf("save address:%v", err.Error())
		}
	} else {
		if err := address.TxRemoveByOrderId(tx, order.Id); err != nil {
			tx.Rollback()
			return fmt.Errorf("save address:%v", err.Error())
		}
	}

	if err := TxSaveMLink(tx, order.Markers, order.Id); err != nil {
		tx.Rollback()
		return fmt.Errorf("save markers:%v", err.Error())
	}
	if err := order.Payments.TxSave(tx, order.Id); err != nil {
		tx.Rollback()
		return fmt.Errorf("save payments:%v", err.Error())
	}

	if err := order.Products.TxSave(tx, order.Id); err != nil {
		tx.Rollback()
		return fmt.Errorf("save product:%v", err.Error())
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	tx = dataBase.MustBegin()
	if err := order.Products.TxSaveModifiers(tx); err != nil {
		return fmt.Errorf("save product:%v", err.Error())
	}
	return tx.Commit()
}
func SaveOfferEvents(order *Order) error {
	dbentity := order.toDBEntity()
	return orders.SaveOfferEvents(&dbentity)
}
func RemoveOrder(id string) error {
	return orders.RemoveById(id)
}
func (o *Order) CalculateCart(priviousVersion *Order, offer *offerentity.RItems) error {
	if priviousVersion != nil && priviousVersion.Paid {
		//если предыдущая версия оплаченна то не производим ни какиз вычеслений и  запрещаем менять корзину

	}
	if o.Paid {

	}
	return nil
}
