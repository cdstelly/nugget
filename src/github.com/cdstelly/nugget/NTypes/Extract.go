package NTypes

import (
	"encoding/json"
	"fmt"
)

type Extract struct {
	PathToExtract string
	AsType string
}

func (m Extract) String() string {
	return string(m.PathToExtract + m.AsType)
}

func (m Extract) JSON() string {
	jsonEncoding, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}