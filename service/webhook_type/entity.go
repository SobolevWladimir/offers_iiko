package webhook_type

type Webhook_type struct {
	Id      string `db:"id" json:"id" valid:"uuid"`
	Alias   string `db:"alias" json:"alias" valid:"-"`
	Name    string `db:"name" json:"name" valid:"-"`
	Comment string `db:"comment" json:"comment" valid:"-"`
	Deleted bool   `db:"deleted" json:"-" valid:"-"`
}
