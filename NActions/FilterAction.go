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

func NewFilterAction(dep BaseAction) FilterAction{
	n := FilterAction{}
	n.dependsOn = dep
	n.executed = false
	return n
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
	if na.dependsOn != nil {
		fmt.Println("filteraction has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.SetFilters(na.filters)

			na.results = na.dependsOn.GetResults()
		}
	}
	// set filters on the parent op
	// execute
	// return results?
//	na.dependsOn.SetFilters(na.filters)
//	operateOn := na.dependsOn.GetResults()

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