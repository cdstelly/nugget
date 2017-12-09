package NTypes

import (
	"encoding/json"
	"fmt"
)

type SHA1 struct {
	Digest string
}

func (s SHA1) String() string {
	return string(s.Digest)
}

func (s SHA1) JSON() string {
	jsonEncoding, err := json.Marshal(s)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}