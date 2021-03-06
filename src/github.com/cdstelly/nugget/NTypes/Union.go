package NTypes

import (
	"fmt"
	"encoding/json"
)

type Union struct {
	Results []string
	AgainstVarName string
}

func (m Union) String() string {
	var ret string
	for _,s:= range m.Results {
		ret += fmt.Sprintf("%s\n", s)
	}
	return ret
}

func (m Union) JSON() string {
	jsonEncoding, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}