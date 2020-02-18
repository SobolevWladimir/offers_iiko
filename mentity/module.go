package mentity

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Name           string
	Label          string
	Actions        Actions
	Route          Route
	Path           string
	Initialization Initialization
}
type Action struct {
	Name     string
	Label    string
	Path     string
	Method   string
	Resource string
	Handler  gin.HandlerFunc
}
type Actions []*Action
type Route func(route *gin.RouterGroup)
type Initialization func()

func (actions *Actions) FindByName(name string) *Action {
	for _, ac := range *actions {
		if ac.Name == name {
			return ac
		}
	}
	return nil
}
func (actions *Actions) FindByPath(path string) *Action {
	for _, ac := range *actions {
		if ac.Path == path {
			return ac
		}
	}
	return nil
}
