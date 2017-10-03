package NActions

import "fmt"
import (
	"../NTypes"
)

type FilterAction struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	results interface{}
}

func (na *FilterAction) BeenExecuted() bool {
	return na.executed
}

func (na *FilterAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *FilterAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}

//todo: dependsOn can be nil..
func (na *FilterAction) Execute() {
	fmt.Println("going to execute filteraction..")
	var myTempResults interface{}
	if na.dependsOn != nil {
		fmt.Println("filteraction has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			//na.dependsOn.SetFilters(na.filters)
			//na.results = na.dependsOn.GetResults()
			myTempResults = na.dependsOn.GetResults()
		}
	}

	//apply the filters
	//does the filter type make sense for the object type?'
	//todo: make DoesPassFilters function a required base function, allow passing in a filter
	var matchedFilter []NTypes.FileInfo
	if files, ok := myTempResults.([]NTypes.FileInfo); ok {
		for _, file := range files {
			if file.DoesPassFilter(na.filters) {
				matchedFilter = append(matchedFilter, file)
			}
		}
	}
	na.results = matchedFilter
	na.executed = true
}

func (na *FilterAction) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *FilterAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}