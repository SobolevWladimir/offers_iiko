package mentity

import (
	"errors"
	"time"
)

type Environment struct {
	Time string `accessfield:"time"`
	Date string `accessfield:"date"`
}

func GetSystemEnviroment() Environment {
	t := time.Now()

	return Environment{
		Time: t.Format("15:04"),
		Date: t.Format("2006-01-02"),
	}

}
func (entity Environment) GetValue(field string) (interface{}, error) {
	switch field {
	case "time":
		return entity.Time, nil
	case "date":
		return entity.Date, nil
	default:
		return nil, errors.New("cant not find field:" + field)
	}
}
