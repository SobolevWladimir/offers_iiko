package server

import (
	"offers_iiko/config"
	"offers_iiko/router"

	"github.com/gin-gonic/gin"
)

var setting config.Core

func SetConfig(conf *config.Config) {
	setting = conf.Core
}
func Run() {
	r := gin.Default()
	router.Route(r)
	r.Run(setting.ServerPort) // listen and serve on 0.0.0.0:8080
}
