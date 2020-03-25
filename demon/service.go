package demon

var demos []DemonInterface

type DemonInterface interface {
	Start()
}

func addDemon(dem DemonInterface) {
	demos = append(demos, dem)
}

func Start() {
	for _, dem := range demos {
		dem.Start()
	}

}
