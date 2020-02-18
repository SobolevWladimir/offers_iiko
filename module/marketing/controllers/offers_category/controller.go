package offers_category

import (
	"altegra_offers/mentity"
	"altegra_offers/service/offers_engine"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
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
	&create,
	&update,
	&remove,
}

var list mentity.Action = mentity.Action{
	Name:   "offers_category_list",
	Label:  "Список категорий акций",
	Path:   "/list",
	Method: "get",
	Handler: func(c *gin.Context) {
		city := c.Query("city")
		city_int, err := strconv.Atoi(city)
		if err != nil {
			c.String(http.StatusBadRequest, "city not number")
			return
		}
		entitys, err := offers_engine.FindByCity(city_int)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, entitys)
	},
}

var create mentity.Action = mentity.Action{
	Name:  "offers_category_create",
	Label: "Создать категорию акции",
	Path:  "/create",
	Handler: func(c *gin.Context) {
		entity := new(offers_engine.Category)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = offers_engine.InsertCategory(entity)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "created")
	},
}

var update mentity.Action = mentity.Action{
	Name:  "offers_category_update",
	Label: "Обновить категорию акции",
	Path:  "/update",
	Handler: func(c *gin.Context) {
		entity := new(offers_engine.Category)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = offers_engine.SaveCategory(entity)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "saved")
	},
}

type RequestRemove struct {
	Id int `json:"id" valid:"uuid"`
}

var remove mentity.Action = mentity.Action{
	Name:  "offers_category_remove",
	Label: "Удалить категорию акции",
	Path:  "/remove",
	Handler: func(c *gin.Context) {
		var req RequestRemove
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		offers_engine.RemoveCategoryById(req.Id)
		c.JSON(200, "removed")
	},
}