package base

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"gopkg.in/guregu/null.v3"
)

type TimeFP time.Time

func (t TimeFP) MarshalJSON() ([]byte, error) {
	stamp := time.Time(t).Format("15:04")
	return json.Marshal(stamp)
}
func (t *TimeFP) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), `"`, "")
	tim, err := time.Parse("15:04", str)
	*t = TimeFP(tim)
	return err
}

type NullTimeFP struct {
	Time  TimeFP
	Valid bool
}

func NewNullTime(tim null.Time) NullTimeFP {
	return NullTimeFP{
		Time:  TimeFP(tim.Time),
		Valid: tim.Valid,
	}
}
func (t *NullTimeFP) ToTime() null.Time {
	return null.Time{
		Time:  time.Time(t.Time),
		Valid: t.Valid,
	}
}
func (t *NullTimeFP) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		err = t.Time.UnmarshalJSON(data)
	case map[string]interface{}:
		ti, tiOK := x["Time"].(string)
		valid, validOK := x["Valid"].(bool)
		if !tiOK || !validOK {
			return fmt.Errorf(`json: unmarshalling object into Go value of type null.Time requires key "Time" to be of type string and key "Valid" to be of type bool; found %T and %T, respectively`, x["Time"], x["Valid"])
		}
		tt := time.Time(t.Time)
		err = tt.UnmarshalText([]byte(ti))
		t.Valid = valid
		return err
	case nil:
		t.Valid = false
		return nil
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type nullTimeFP", reflect.TypeOf(v).Name())
	}
	t.Valid = err == nil
	return err
}
func (t NullTimeFP) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

func (t NullTimeFP) MarshalText() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}
