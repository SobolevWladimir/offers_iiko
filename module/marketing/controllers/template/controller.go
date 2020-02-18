package template

import (
	"altegra_offers/mentity"
	"altegra_offers/service/offers_engine"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const CONTROLLER_NAME = "offers_template"

func FindActionByPath(path string) *mentity.Action {
	return Actions.FindByPath(path)
}

var Entity mentity.Controller = mentity.Controller{
	Name:    CONTROLLER_NAME,
	Label:   "Шаблоны",
	Path:    "/template",
	Actions: Actions,
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
	&create,
	&update,
	&remove,
}

var list mentity.Action = mentity.Action{
	Name:   "offers_template_list",
	Label:  "Список шаблонов",
	Path:   "/list",
	Method: "get",
	Handler: func(c *gin.Context) {
		entitys, err := offers_engine.FindAllTemplates()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusOK, entitys)
	},
}

var create mentity.Action = mentity.Action{
	Name:  "offers_template_create",
	Label: "Создать Шаблон",
	Path:  "/create",
	Handler: func(c *gin.Context) {
		entity := new(offers_engine.Template)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = entity.Insert()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "created")
	},
}

var update mentity.Action = mentity.Action{
	Name:  "offers_template_update",
	Label: "Обновить категорию",
	Path:  "/update",
	Handler: func(c *gin.Context) {
		entity := new(offers_engine.Template)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = entity.Save()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "saved")
	},
}

type RequestRemove struct {
	Id string `json:"id" valid:"uuid"`
}

var remove mentity.Action = mentity.Action{
	Name:  "offers_template_remove",
	Label: "Удалить шаблон",
	Path:  "/remove",
	Handler: func(c *gin.Context) {
		var req RequestRemove
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		offers_engine.RemoveTemplateById(req.Id)
		c.JSON(200, "removed")
	},
}
