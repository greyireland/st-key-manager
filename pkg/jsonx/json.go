package jsonx

import (
	"encoding/json"
)

func JSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
