package util

import (
	"strconv"
	"encoding/json"
)

type Str string

func (str *Str) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*str = Str(s)
		return nil
	}
	*str = Str(data)
	return nil
}

func (str Str) Str() string {
	return string(str)
}

func (str Str) GetStrInt() int {
	i, _ := strconv.Atoi(str.Str())
	return i
}
