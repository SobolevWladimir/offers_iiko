package coupon

import (
	"altegra_offers/mentity"
	"altegra_offers/service/coupon"
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const CONTROLLER_NAME = "coupon"

func FindActionByPath(path string) *mentity.Action {
	return Actions.FindByPath(path)
}

var Entity mentity.Controller = mentity.Controller{
	Name:    CONTROLLER_NAME,
	Label:   "Купоны",
	Path:    "/coupon",
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
	&generate,
}

var list mentity.Action = mentity.Action{
	Name:   "coupon_list",
	Label:  "Список купонов",
	Path:   "/list",
	Method: "get",
	Handler: func(c *gin.Context) {
		category := c.Query("category")
		cat_int, err := strconv.Atoi(category)
		if err != nil {
			c.String(http.StatusBadRequest, "category is  number")
			return
		}
		entitys, err := coupon.FindByCategory(cat_int)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusOK, entitys)
	},
}
var create mentity.Action = mentity.Action{
	Name:  "coupon_create",
	Label: "Создать купон",
	Path:  "/create",
	Handler: func(c *gin.Context) {
		entity := new(coupon.Coupon)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		exist, err := CheckIsExist(entity)
		if err != nil {
			c.String(http.StatusBadRequest, "error check for exist. error:"+err.Error())
			return
		}
		if exist {
			c.String(http.StatusBadRequest, " coupon already is exist")
			return
		}
		err = coupon.Insert(entity)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "created")
	},
}
var update mentity.Action = mentity.Action{
	Name:  "coupon_update",
	Label: "Обновить купон",
	Path:  "/update",
	Handler: func(c *gin.Context) {
		entity := new(coupon.Coupon)
		if err := c.ShouldBindJSON(entity); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		valid, err := govalidator.ValidateStruct(entity)
		if !valid || err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = coupon.Save(entity)
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
	Name:  "coupon_remove",
	Label: "Удалить купон",
	Path:  "/remove",
	Handler: func(c *gin.Context) {
		var req RequestRemove
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		coupon.RemoveById(req.Id)
		c.JSON(200, "removed")
	},
}
var generate mentity.Action = mentity.Action{
	Name:  "coupon_generate",
	Label: "Генерация купонов",
	Path:  "/generate",
	Handler: func(c *gin.Context) {
		var gdata GData
		if err := c.ShouldBindJSON(&gdata); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		coupons, err := gdata.Generate()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		coupons, err = FilterNotExist(&coupons, gdata.Category)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = coupon.InsertCoupons(&coupons)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.String(200, "generate")
	},
}
