package router

import (
	"offers_iiko/config"
	"offers_iiko/router/middleware"

	"github.com/gin-gonic/gin"
)

var modules []ModuleInterface

type ModuleInterface interface {
	RelativePath() string
	SetConfig(config.Core)
	Route(route *gin.RouterGroup)
}

// отправить в плагин   который проверяет авторизацию
func Route(route *gin.Engine) {

	route.Use(middleware.Handler())
	//route.GET("/ping", testGet)
	for _, module := range modules {
		module.Route(route.Group("/" + module.RelativePath()))

	}

}
func addModule(module ModuleInterface) {
	modules = append(modules, module)
}
