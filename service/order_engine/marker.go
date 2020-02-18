package order_engine

import (
	"altegra_offers/service/order_engine/markers"
)

type Marker string
type Markers []Marker

func ConventMarkers(mas []string) Markers {
	result := Markers{}
	for _, m := range mas {
		result = append(result, Marker(m))
	}
	return result
}
func FindAllMarders() (Markers, error) {
	mars, err := markers.FindAll()
	return ConventMarkers(mars), err
}

func InsertMarker(name string) error {
	isset, err := markers.ExistByName(name)
	if err != nil {
		return err
	}
	if isset {
		return markers.Save(name)
	}
	return markers.Insert(name)
}
func RemoveMarker(name string) error {
	return markers.RemoveByName(name)
}
