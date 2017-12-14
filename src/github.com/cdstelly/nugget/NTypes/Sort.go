package NTypes

import (
	"encoding/json"
	"fmt"
)

type Sort struct {
	Field string
}


func (s Sort) String() string {
	return string(s.Field)
}

func (s Sort) JSON() string {
	jsonEncoding, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}