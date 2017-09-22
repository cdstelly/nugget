package NActions

import "fmt"
import (
	"../NTypes"
	"sort"
	"crypto/md5"
)

type SortAction struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	results interface{}
}

func (na *SortAction) BeenExecuted() bool {
	return na.executed
}

func (na *SortAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *SortAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}
func (na *SortAction) Execute() {
	if na.dependsOn != nil {
		fmt.Println("sortaction has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	fmt.Println("going to execute sort..")

	operateOn := na.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes. ); ok {


	}
	na.executed = true
}

func (na *SortAction) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *SortAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}