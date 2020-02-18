package router

import (
	"offers_iiko/module/marketing"
)

func init() {
	//	addModule(address.New())
	//	addModule(settings.New())
	//	addModule(staff.New())
	//	addModule(sync.New())
	//	addModule(accessm.New())
	//	addModule(organization.New())
	//	addModule(menu.New())
	//	addModule(uploads.New())
	//	addModule(sales.New())
	addModule(marketing.New())
	//	addModule(pricelist.New())
	//	addModule(events.New())
}
