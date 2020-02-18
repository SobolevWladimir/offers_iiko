package base

import (
	"encoding/json"
	"strings"
	"time"
)

//Дата в
type DateFP time.Time

func (t DateFP) MarshalJSON() ([]byte, error) {
	stamp := time.Time(t).Format("2006-01-02")
	return json.Marshal(stamp)
}
func (t *DateFP) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), `"`, "")
	tim, err := time.Parse("2006-01-02", str)
	*t = DateFP(tim)
	return err
}
