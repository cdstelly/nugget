package NTypes

import (
	"encoding/json"
	"fmt"
)

type ProcessList struct {
	Processes string
}

func (pl ProcessList) String() string {
	return string(pl.Processes)
}

func (pl ProcessList) JSON() string {
	jsonEncoding, err := json.Marshal(pl)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}