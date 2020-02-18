package order_engine

import (
	"altegra_offers/service/order_engine/payment"

	"github.com/jmoiron/sqlx"
)

type Payment struct {
	Id     string  `json:"id"`
	Type   int     `json:"type"`
	Amount float32 `json:"amount"`
}
type Payments []Payment

func toPayment(p *payment.Payment) Payment {
	result := Payment{}
	result.Id = p.Id
	result.Type = p.Type
	result.Amount = p.Amount
	return result
}
func (p *Payment) toDBEntity(order string) payment.Payment {
	result := payment.Payment{}
	result.Id = p.Id
	result.Type = p.Type
	result.Amount = p.Amount
	result.Order = order
	return result
}
func toPayments(ps *payment.Payments) Payments {
	result := Payments{}
	for _, pay := range *ps {
		result = append(result, toPayment(&pay))
	}
	return result
}
func (ps *Payments) toDBEntity(order string) payment.Payments {
	result := payment.Payments{}
	for _, pay := range *ps {
		result = append(result, pay.toDBEntity(order))
	}
	return result
}
func FindAllPaymentsByOrders(order string) (Payments, error) {
	pays, err := payment.FindAllByOrder(order)
	return toPayments(&pays), err
}
func (ps *Payments) Insert(order string) error {
	pay := ps.toDBEntity(order)
	return payment.InsertArray(&pay)
}
func (ps *Payments) Save(order string) error {
	pay := ps.toDBEntity(order)
	if err := payment.RemoveNotExist(&pay, order); err != nil {
		return err
	}
	return payment.SaveExist(&pay)

}
func (ps *Payments) TxInsert(tx *sqlx.Tx, order string) error {
	pay := ps.toDBEntity(order)
	return payment.TxInsertArray(tx, &pay)
}
func (ps *Payments) TxSave(tx *sqlx.Tx, order string) error {
	pay := ps.toDBEntity(order)
	if err := payment.TxRemoveNotExist(tx, &pay, order); err != nil {
		return err
	}
	return payment.TxSaveExist(tx, &pay)

}
