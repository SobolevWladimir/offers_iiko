package middleware

import (
	"offers_iiko/mentity"

	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		given := mentity.Given{}
		c.Set("given", given)
		c.Next()
	}
}
