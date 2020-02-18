package order_engine

import (
	"encoding/json"
)

type NullAddress struct {
	Address Address
	Valid   bool
}

func (t NullAddress) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Address)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string, object (e.g. pq.NullTime and friends)
// and null input.
func (t *NullAddress) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v == nil {
		t.Valid = false
		return nil
	}
	err = json.Unmarshal(data, &t.Address)
	t.Valid = err == nil
	return err
}
