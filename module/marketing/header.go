package marketing

import (
	"altegra_offers/config"
	"altegra_offers/mentity"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

var setting config.Core
var controllers []*mentity.Controller

type ModuleObject struct {
}

func New() *ModuleObject {
	return new(ModuleObject)
}

func (m *ModuleObject) SetConfig(sett config.Core) {
	setting = sett
}
func (m *ModuleObject) Route(route *gin.RouterGroup) {
	checkControllers()
	route.Use(Handler())
	for _, controller := range controllers {
		if controller.Initialization != nil {
			controller.Initialization()
		}
		controller.Route(route.Group("/" + controller.Path))

	}
}
func addControler(con *mentity.Controller) {
	controllers = append(controllers, con)
}
func checkControllers() {
	contr := make(map[string]string)
	actions := make(map[string]string)
	for _, con := range controllers {
		if clabel, ok := contr[con.Name]; ok {

			log.Fatal(" module: " + reflect.TypeOf(ModuleObject{}).PkgPath() + ". Не могу добавить контролер Name: " + con.Name + " Label:" + clabel + " \n Этот контролер уже был проинициализирован ранее")
		} else {
			contr[con.Name] = con.Label
		}
		for _, ac := range con.Actions {
			if alabel, aok := actions[ac.Name]; aok {

				log.Fatal(" module: " + reflect.TypeOf(ModuleObject{}).PkgPath() + ". Не могу добавить action Name: " + ac.Name + " Label:" + alabel + " \n Эта action уже была проинициализированa ранее")

			} else {
				actions[ac.Name] = ac.Label
			}
		}
	}
}
