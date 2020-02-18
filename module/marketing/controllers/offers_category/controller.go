package offers_category

import (
	"net/http"
	"offers_iiko/mentity"
	"strings"

	"github.com/gin-gonic/gin"
)

const CONTROLLER_NAME = "offers_category"

func FindActionByPath(path string) *mentity.Action {
	return Actions.FindByPath(path)
}

var Entity mentity.Controller = mentity.Controller{
	Name:    CONTROLLER_NAME,
	Label:   "Категории акций",
	Path:    "/offers_category",
	Actions: Actions,
	Initialization: func() {
	},
	Route: func(route *gin.RouterGroup) {
		route.Use(func(c *gin.Context) {
			path := strings.Replace(c.Request.URL.Path, route.BasePath(), "", 1)
			ac := FindActionByPath(path)
			if ac != nil {
				given := c.MustGet("given").(mentity.Given)
				given.Action = ac.Name
				given.Controller = CONTROLLER_NAME
				c.Set("given", given)
			}
		})
		for _, ac := range Actions {
			switch ac.Method {
			case "get":
				route.GET(ac.Path, ac.Handler)
			default:
				route.POST(ac.Path, ac.Handler)

			}
		}
	},
}
var Actions mentity.Actions = mentity.Actions{
	&list,
}

var list mentity.Action = mentity.Action{
	Name:   "offers_category_list",
	Label:  "Список категорий акций",
	Path:   "/list",
	Method: "get",
	Handler: func(c *gin.Context) {
		c.String(http.StatusOK, "of")
	},
}
