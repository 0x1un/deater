package utils

import "encoding/json"

func IsJsonAble(b []byte) interface{} {
	var r interface{}
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}
	return r
}
