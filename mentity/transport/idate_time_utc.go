package transport

import (
	"encoding/json"
	"time"
)

type IDateTimeUTC time.Time

func (t IDateTimeUTC) MarshalJSON() ([]byte, error) {
	time := time.Time(t)
	result := time.Format("2006-01-02T15:04:05Z07:00")
	return json.Marshal(result)
}
