package offerentity

type Action struct {
	Type int         `json:"type"` //	 тип действия
	Data interface{} `json:"data"` // Данные действия
}
type Actions []Action

//типы действий
const (
	TypePresent              = 1
	TypeSpecialDiscount      = 2
	TypeDiscount             = 3
	TypeUpsale               = 6 //	 подсказка
	TypeActionBonuses        = 7
	TypeActionClientCategory = 8
	TypeActionCoupon         = 1000
)

func (as *Actions) ContainsiByType(t int) bool {
	for _, a := range *as {
		if a.Type == t {
			return true
		}
	}
	return false
}
