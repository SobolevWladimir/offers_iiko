package setting

import "gopkg.in/guregu/null.v3"

type SitySetting struct {
	ID       int
	SityID   int
	SityCode string
	Info     string
	Phone    string
}
type IikoSetting struct {
	ID           int         `db:"id"`
	SityId       int         `db:"sity_id"`
	UserID       null.String `db:"user_id"`
	UserSecret   null.String `db:"user_secret"`
	Organization null.String `db:"organization"`
	CityId       null.String `db:"cityId"`
	TerminalId   null.String `db:"terminalid"`
}
type IikoSettings []IikoSetting
