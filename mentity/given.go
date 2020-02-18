package mentity

import (
	"errors"
)

const (
	TypeEnvironment int = 0
	TypeSubject     int = 1
	TypeResource    int = 2
)

// Отпровляемое значение
type Given struct {
	Module      string
	Action      string
	Controller  string
	Subject     ObjectInterface
	Resource    ObjectInterface
	Environment ObjectInterface
}

func (g *Given) GetValue(typ int, field string) (interface{}, error) {
	var ot ObjectInterface
	switch typ {
	case TypeEnvironment:
		ot = g.Environment
	case TypeSubject:
		ot = g.Subject
	case TypeResource:
		ot = g.Resource
	default:
		ot = nil
	}
	if ot != nil {
		return ot.GetValue(field)
	}
	return nil, errors.New("I can not determine the type in fuction given or inteface is nil")
}

type ObjectInterface interface {
	GetValue(field string) (interface{}, error)
}
