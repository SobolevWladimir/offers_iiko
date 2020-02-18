package products

import (
	"gopkg.in/guregu/null.v3"
)

type Product struct {
	Id       string      `db:"id"`
	Order    string      `db:"order"`
	Product  string      `db:"product"`
	Quantity float32     `db:"quantity"`
	Comment  null.String `db:"comment"`
}
type Products []Product

func (ps *Products) getIds() []string {
	result := []string{}
	for _, p := range *ps {
		if len(p.Id) > 0 {
			result = append(result, p.Id)
		}
	}
	return result
}
