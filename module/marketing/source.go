package marketing

import (
	"altegra_offers/mentity"
	"altegra_offers/module/marketing/controllers/coupon"
	"altegra_offers/module/marketing/controllers/coupon_category"
	"altegra_offers/module/marketing/controllers/offer"
	"altegra_offers/module/marketing/controllers/offers_category"
	"altegra_offers/module/marketing/controllers/template"

	"github.com/gin-gonic/gin"
)

const ModuleName = "marketing"

func (m *ModuleObject) RelativePath() string {

	return "/" + ModuleName
}

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		given := c.MustGet("given").(mentity.Given)
		given.Module = ModuleName
		c.Set("given", given)
		c.Next()
	}
}
func init() {
	initControllers()
}

// регистрируем котролеры сдесь
func initControllers() {
	addControler(&coupon_category.Entity)
	addControler(&coupon.Entity)
	addControler(&offers_category.Entity)
	addControler(&offer.Entity)
	addControler(&template.Entity)
}
