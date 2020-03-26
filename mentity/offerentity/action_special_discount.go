package offerentity

type ActionSpecialDiscount struct {
	Max         int     `json:"max"`
	SaleType    int     `json:"sale_type"` // 0-%, 1-сумма 2-фикс
	Target      int     `json:"target"`
	Value       float32 `json:"value"`
	OrderItemId string  `json:"orderItemId"` // id  позиции в заказе
}
type ActionsSpecialDiscount []ActionSpecialDiscount
type ActionSpecialDiscountTarget []string

type TypeValueSpecialDiscount int

const (
	TypeDiscountCategory TypeValueSpecialDiscount = 0
	TypeDiscountGoods    TypeValueSpecialDiscount = 1
	TypeDiscountDish     TypeValueSpecialDiscount = 2
	TypeDiscountPrepack  TypeValueSpecialDiscount = 3
)

func (a *ActionSpecialDiscountTarget) Contains(target string) bool {
	for _, item := range *a {
		if target == item {
			return true
		}
	}
	return false
}
