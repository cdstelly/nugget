package NTypes

import "fmt"

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
