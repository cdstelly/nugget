package NTypes

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type HTTP struct {
	Host string
	Method string
	Length int
	Data []byte
	RawRequest http.Request
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