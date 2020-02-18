package webhook

type Webhook struct {
	Id      string `db:"id" json:"id" valid:"uuid"`
	Type    string `db:"type" json:"type" valid:"uuid"`
	URL     string `db:"url" json:"url" valid:"-"`
	Deleted bool   `db:"deleted" json:"-" valid:"-"`
}
