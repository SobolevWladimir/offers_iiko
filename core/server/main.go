package server

import (
	"fmt"
	"offers_iiko/config"
	"offers_iiko/router"

	"github.com/gin-gonic/gin"
)

var setting config.Core

func SetConfig(conf *config.Config) {
	setting = conf.Core
}
func Run(mode int) {
	if mode == config.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	router.Route(r)
	fmt.Println("Сервер запущен! порт: ", setting.ServerPort)
	r.Run(setting.ServerPort) // listen and serve on 0.0.0.0:8080
}
