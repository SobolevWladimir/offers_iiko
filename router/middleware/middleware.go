package middleware

import (
	"altegra_offers/mentity"

	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		given := mentity.Given{}
		c.Set("given", given)
		c.Next()
	}
}
