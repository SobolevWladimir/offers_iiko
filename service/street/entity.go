package street

type Street struct {
	ID        string `db:"id" json:"id" valid:"uuid"`
	Name      string `db:"name" json:"name" valid:"utfletternumspace"`
	City      string `db:"city" json:"city" valid:"uuid"`
	Type      string `db:"type" json:"type" valid:"utfletternumspace"`
	TypeShort string `db:"type_short" json:"type_short" valid:"utfletternumspace"`
	Deleted   bool   `db:"deleted" json:"deleted" valid:"-"`
}
type Streets []Street

func (s *Streets) FilterByCity(city string) Streets {
	result := Streets{}
	for _, st := range *s {
		if st.City == city {
			result = append(result, st)
		}
	}
	return result
}
