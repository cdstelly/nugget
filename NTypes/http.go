package NTypes

import (
	"encoding/json"
	"fmt"
)


type HTTP struct {
	Host string
	Data []byte
}

func (m HTTP) String() string {
	return string(m.Host)
}

func (m HTTP) JSON() string {
	jsonEncoding, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}