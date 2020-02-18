package payment

type Payment struct {
	Id     string  `db:"id"`
	Order  string  `db:"order"`
	Type   int     `db:"type"`
	Amount float32 `db:"amount"`
	//@fixme add isdone
}

type Payments []Payment

func (ps *Payments) getIds() []string {
	result := []string{}
	for _, p := range *ps {
		if len(p.Id) > 0 {
			result = append(result, p.Id)
		}
	}
	return result
}
