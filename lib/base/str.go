package base

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Обычная стринга в json  может быть  int
type StringInt string

func (str *StringInt) UnmarshalJSON(data []byte) error {
	var v interface{}
	var err error
	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	switch x := v.(type) {
	case string:
		*str = StringInt(x)
	case float64:
		fmt.Println("float64", x, strconv.FormatFloat(x, 'E', -1, 64))
		*str = StringInt(strconv.FormatFloat(x, 'f', -1, 64))
	default:
		err = fmt.Errorf("json: cannot unmarshal %v into Go value of type null.String", reflect.TypeOf(v).Name())
	}
	return err
}
