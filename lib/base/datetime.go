package base

import (
	"encoding/json"
	"strings"
	"time"
)

type DateTimeFP time.Time

func (t DateTimeFP) MarshalJSON() ([]byte, error) {
	stamp := time.Time(t).Format("2006-01-02 15:04:05")
	return json.Marshal(stamp)
}
func (t *DateTimeFP) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), `"`, "")
	tim, err := time.Parse("2006-01-02 15:04:05", str)
	*t = DateTimeFP(tim)
	return err
}
