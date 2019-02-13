package NTypes

import (
	"encoding/json"
	"fmt"
	"time"
)

type Datetime struct {
	theDatetime time.Time
}

func (m Datetime) String() string {
	return string(m.String())
}

func (m Datetime) JSON() string {
	jsonEncoding, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}
