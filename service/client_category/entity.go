package client_category

type ClientCategory struct {
	Id      string `db:"id" json:"id" valid:"uuid"`
	Name    string `db:"name" json:"name" valid:"required" accessfield:"name"`
	Deleted bool   `db:"deleted" json:"-" valid:"-"`
}
type ClientCategorys []ClientCategory
