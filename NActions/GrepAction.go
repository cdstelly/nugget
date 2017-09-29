package NActions

import (
	"../NTypes"
	"fmt"
	"regexp"
	"reflect"
	"strings"
)

type GrepAction struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter


	Expression string
	results []string
}

func (na *GrepAction) BeenExecuted() bool {
	return na.executed
}

func (na *GrepAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *GrepAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (na *GrepAction) Execute() {
	if na.dependsOn != nil {
		fmt.Println("grep action has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	fmt.Println("going to execute grep..")

	if na.dependsOn == nil {
		fmt.Println("Error! Was not able to compute SHA1 because dependency is nil")
		return
	}

	var operateOn []string

	givenData := na.dependsOn.GetResults()
	var validStr = regexp.MustCompile(strings.Trim(na.Expression, `"`))

	//if we're given something other than strings, get the string representation of it:
	if opr, ok := givenData.([]NTypes.NPacket); ok {
		for _,p := range opr {
			if len(p.String()) > 1 {
				operateOn = append(operateOn, p.String())
			}
		}
	} else {
		fmt.Println("Error: grep was unable to operate on type: ", reflect.TypeOf(operateOn))
	}

	var mystrarr []string
	//todo: quite possibly, the hackiest code ever. first priority after proof of concept is to take another look at how we pass around objects to actions
	var validHost = regexp.MustCompile(`^Host.*`)
	for _, inString := range operateOn {
		var getReq map[string]string
		getReq = make(map[string]string)

		for _, s := range strings.Split(inString,"\n") {
			if validStr.MatchString(s) {
				getReq["URI"] = strings.TrimSpace(s)
				//na.results = append(na.results, s)
			}
			if validHost.MatchString(s) {
				//na.results = append(na.results, s)
				getReq["Host"] = strings.TrimSpace(s)
			}
		}

		if val, ok := getReq["Host"]; ok {
			if val2, ok2 := getReq["URI"]; ok2 {
				mystrarr = append(mystrarr, val+" "+val2 )
			}
		}

	}
	na.results = mystrarr
	na.executed = true
}

func (na *GrepAction) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *GrepAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}