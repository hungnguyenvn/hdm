package hdm

import (
	"encoding/json"
	"strconv"
)

type FloatString float64

func (ft *FloatString) UnmarshalJSON(data []byte) error {
	if data[0] != '"' {
		return json.Unmarshal(data, (*float64)(ft))
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if len(s) == 0 {
		return nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*ft = FloatString(f)
	return nil
}
