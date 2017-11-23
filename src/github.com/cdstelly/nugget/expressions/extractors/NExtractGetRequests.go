package extractors

import (
	"github.com/cdstelly/nugget/NTypes"
	"github.com/cdstelly/nugget/expressions/transforms"
)

type NExtractGetRequests struct {
	executed  bool
	dependsOn expressions.BaseAction
	filters []NTypes.Filter

	results []string
}

func (na *NExtractGetRequests) BeenExecuted() bool {
	return na.executed
}

func (na *NExtractGetRequests) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *NExtractGetRequests) SetDependency(action expressions.BaseAction) {
	na.dependsOn = action
}
func (na *NExtractGetRequests) Execute() {
	if na.dependsOn != nil {
		//fmt.Println("getGetReqs has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	//fmt.Println("going to execute getGetReqs..")

	operateOn := na.dependsOn.GetResults()

	if _, ok := operateOn.([]NTypes.NPacket); ok {



	}

	na.executed = true
}

func (na *NExtractGetRequests) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *NExtractGetRequests) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}