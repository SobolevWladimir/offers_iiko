package offer

import (
	"altegra_offers/eoffer"
	"altegra_offers/mentity"
	"altegra_offers/mentity/offerentity"
	"altegra_offers/service/offers_engine"
	"altegra_offers/service/order_engine"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const CONTROLLER_NAME = "offer"

func FindActionByPath(path string) *mentity.Action {
	return Actions.FindByPath(path)
}

var Entity mentity.Controller = mentity.Controller{
	Name:    CONTROLLER_NAME,
	Label:   "Акции",
	Path:    "/offer",
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
	&storage,
	&calculate,
}

var list mentity.Action = mentity.Action{
	Name:   "offer_list",
	Label:  "Список  Акций",
	Path:   "/list",
	Method: "get",
	Handler: func(c *gin.Context) {
		category := c.Query("category")
		cat_int, err := strconv.Atoi(category)
		if err != nil {
			c.String(http.StatusBadRequest, "category is not  number")
			return
		}
		entitys, err := offers_engine.FindOffersByCategory(cat_int)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		//filter
		c.JSON(http.StatusOK, entitys)
	},
}

var create mentity.Action = mentity.Action{
	Name:  "offers_create",
	Label: "Создать Акцию",
	Path:  "/create",
	Handler: func(c *gin.Context) {
		entity := new(offers_engine.Policy)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, "to json", err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = offers_engine.InsertPolicy(entity)
		if err != nil {
			c.String(http.StatusBadRequest, "db:"+err.Error())
			return
		}
		c.String(http.StatusOK, "created")
	},
}

var update mentity.Action = mentity.Action{
	Name:  "offers_update",
	Label: "Обновить  Акцию",
	Path:  "/update",
	Handler: func(c *gin.Context) {
		entity := new(offers_engine.Policy)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = offers_engine.SavePolicy(entity)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "saved")
	},
}

type RequestRemove struct {
	Id int `json:"id" `
}

var remove mentity.Action = mentity.Action{
	Name:  "offers_remove",
	Label: "Удалить Акцию",
	Path:  "/remove",
	Handler: func(c *gin.Context) {
		var req RequestRemove
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		offers_engine.RemovePolicyById(req.Id)
		c.JSON(200, "removed")
	},
}
var storage mentity.Action = mentity.Action{
	Name:   "offers_storage",
	Label:  "Переменные среды",
	Path:   "/storage",
	Method: "get",
	Handler: func(c *gin.Context) {
		city := c.Query("city")
		city_int, err := strconv.Atoi(city)
		if err != nil {
			c.String(http.StatusBadRequest, "city is not number")
		}

		storage := offers_engine.GetStandartStorage(city_int)
		c.JSON(http.StatusOK, &storage)

	},
}
var calculate mentity.Action = mentity.Action{
	Name:  "offers_calculate",
	Label: "Расчитаать акции для заказа",
	Path:  "/calculate",
	Handler: func(c *gin.Context) {
		entity := new(order_engine.Order)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		oldEntity, err := order_engine.FindOrderById(entity.Id)
		var errc error
		var offer offerentity.Allegiance
		if err != nil {
			offer, errc = eoffer.Calculate(entity, nil)
		} else {
			offer, errc = eoffer.Calculate(entity, &oldEntity)
		}
		if errc != nil {
			c.String(http.StatusBadRequest, errc.Error())
			return
		}
		c.JSON(http.StatusOK, offer)
	},
}
