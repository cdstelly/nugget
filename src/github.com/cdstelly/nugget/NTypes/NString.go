package NTypes

import (
	"encoding/json"
	"fmt"
)
//todo remove this type.
type NString struct {
	Text string
}
func (p NString) String() string {
	return p.Text
}

func (p NString) JSON() string {
	jsonEncoding, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}