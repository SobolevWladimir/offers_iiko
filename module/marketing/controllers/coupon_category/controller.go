package coupon_category

import (
	"altegra_offers/mentity"
	"altegra_offers/service/coupon_category"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const CONTROLLER_NAME = "coupon_category"

func FindActionByPath(path string) *mentity.Action {
	return Actions.FindByPath(path)
}

var Entity mentity.Controller = mentity.Controller{
	Name:    CONTROLLER_NAME,
	Label:   "Категории купонов",
	Path:    "/coupon_category",
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
	Name:   "coupon_category_list",
	Label:  "Список категорий купонов",
	Path:   "/list",
	Method: "get",
	Handler: func(c *gin.Context) {
		city := c.Query("city")
		if len(city) == 0 {
			c.String(http.StatusBadRequest, "city can't be  null")
			return
		}

		city_int, err := strconv.Atoi(city)
		if err != nil {
			c.String(http.StatusBadRequest, "city is not number")
			return
		}
		entitys, err := coupon_category.FindByCity(city_int)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusOK, entitys)
	},
}
var create mentity.Action = mentity.Action{
	Name:  "coupon_category_create",
	Label: "Создать категорию купонов",
	Path:  "/create",
	Handler: func(c *gin.Context) {
		entity := new(coupon_category.Category)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = coupon_category.Insert(entity)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "created")
	},
}
var update mentity.Action = mentity.Action{
	Name:  "coupon_category_update",
	Label: "Обновить категорию  купонов",
	Path:  "/update",
	Handler: func(c *gin.Context) {
		entity := new(coupon_category.Category)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = coupon_category.Save(entity)
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
	Name:  "coupon_category_remove",
	Label: "Удалить категорию купонов",
	Path:  "/remove",
	Handler: func(c *gin.Context) {
		var req RequestRemove
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		coupon_category.RemoveById(req.Id)
		c.JSON(200, "removed")
	},
}
