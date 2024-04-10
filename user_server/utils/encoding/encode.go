package encoding

import (
	"encoding/json"
)

func EncodeIgnoreError(data interface{}) string {
	byteSlice, _ := json.Marshal(data)
	return string(byteSlice)
}
