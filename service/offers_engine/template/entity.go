package template

type Template struct {
	Id      string `db:"id"`
	SetRule string `db:"setrule"`
	Sort    int    `db:"sort"`
}
type Templates []Template
