package NTypes

import (
	"encoding/json"
	"fmt"
)

type SHA256 struct {
	Digest string
}

func (s SHA256) String() string {
	return string(s.Digest)
}

func (s SHA256) JSON() string {
	jsonEncoding, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}