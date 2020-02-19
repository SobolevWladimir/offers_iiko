package marketing

import (
	"offers_iiko/mentity"
	"offers_iiko/module/marketing/controllers/offers"

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
	addControler(&offers.Entity)
}
