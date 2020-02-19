package offers

import (
	"net/http"
	"offers_iiko/mentity"
	"offers_iiko/mentity/transport"
	"strings"

	"github.com/gin-gonic/gin"
)

const CONTROLLER_NAME = "offers_category"

func FindActionByPath(path string) *mentity.Action {
	return Actions.FindByPath(path)
}

var Entity mentity.Controller = mentity.Controller{
	Name:    CONTROLLER_NAME,
	Label:   "акций",
	Path:    "/offers",
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
	&check,
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
var check mentity.Action = mentity.Action{
	Name:  "offer_check",
	Label: "проверить акции",
	Path:  "/check",
	Handler: func(c *gin.Context) {
		entity := new(transport.AOrderRequest)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, entity)
	},
}
