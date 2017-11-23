package extractors

import (
	"github.com/cdstelly/nugget/NTypes"
	"github.com/cdstelly/nugget/expressions/transforms"
)

type ExtractMemory struct {
	executed  bool
	dependsOn expressions.BaseAction
	filters []NTypes.Filter

	MemoryLocation string

}

func (na *ExtractMemory) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractMemory) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractMemory) SetDependency(action expressions.BaseAction) {
	na.dependsOn = action
}

func (na *ExtractMemory) Execute() {
	na.executed = true
}

func (na *ExtractMemory) GetResults() interface{}{
	if na.executed == false {
		na.Execute()
	}
	return na.MemoryLocation
}

func (na *ExtractMemory) SetFilters(filters []NTypes.Filter) {
	na.filters = filters
}