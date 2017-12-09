package NTypes

import (
	"encoding/json"
	"fmt"
)

type Filter struct {
	Field string
	Op string
	Value string
}

func (f Filter) String() string {
	return string(f.Field + f.Op + f.Value)
}

func (f Filter) JSON() string {
	jsonEncoding, err := json.Marshal(f)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}
