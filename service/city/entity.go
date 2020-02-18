package city

type City struct {
	ID      string `db:"id" json:"id" valid:"uuid"`
	Name    string `db:"name" json:"name" valid:"utfletternumspace" accessfield:"name"`
	Deleted bool   `db:"deleted" json:"deleted" valid:"-"`
}
type Citys []City
