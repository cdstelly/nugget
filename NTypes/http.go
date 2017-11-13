package NTypes

import (
	"encoding/json"
	"fmt"
)

type HTTPHeaders struct {
	Method string
	Host string
	UserAgent string
	//Accept string
	//AcceptLanguage string
	//AcceptEncoding string
	//AcceptCharset string
	//KeepAlive string
	//Connection string
	Referer string
	Cookie string
}

type HTTP struct {
	Headers HTTPHeaders
	Host string
	Data string
}

func (m HTTP) String() string {
	return string(m.Host + m.Data)
}

func (m HTTP) JSON() string {
	jsonEncoding, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}