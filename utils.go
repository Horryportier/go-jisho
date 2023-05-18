package gojisho

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ResolveUnmarshalErr(data []byte, err error) string {
	if e, ok := err.(*json.UnmarshalTypeError); ok {
		// grab stuff ahead of the error
		var i int
		for i = int(e.Offset) - 1; i != -1 && data[i] != '\n' && data[i] != ','; i-- {
		}
		info :=  strings.TrimSpace(string(data[i+1 : int(e.Offset)]))
		s := fmt.Sprintf("%s - at: %s", e.Error(), info)
		return s
	}
	if e, ok := err.(*json.UnmarshalFieldError); ok {
		return e.Error()
	}
	if e, ok := err.(*json.InvalidUnmarshalError); ok {
		return e.Error()
	}
	return err.Error()
}
