package NTypes

import (
	"encoding/json"
	"fmt"
)

type MD5 struct {
	Digest string
	HashOf interface{}
}

func (m MD5) String() string {
	return string(m.Digest)
}

func (m MD5) JSON() string {
	jsonEncoding, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}