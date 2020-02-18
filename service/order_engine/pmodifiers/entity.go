package pmodifiers

type Modifier struct {
	Id       string  `db:"id"`
	Product  string  `db:"product"`
	Modifier string  `db:"modifier"`
	Quantity float32 `db:"quantity"`
}
type Modifiers []Modifier

func (ps *Modifiers) getIds() []string {
	result := []string{}
	for _, p := range *ps {
		if len(p.Id) > 0 {
			result = append(result, p.Id)
		}
	}
	return result
}
