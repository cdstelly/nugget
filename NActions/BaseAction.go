package NActions

import "../NTypes"

type BaseAction interface {
	BeenExecuted() bool
	DependencySatisfied() bool
	SetDependency(action BaseAction)
	Execute()
	GetResults() interface{}
	SetFilters([]NTypes.Filter)

	// ExternalPassesFilterCheck([]NTypes.Filter) bool idea with this would be to cull results from things it depends on, saving execution cycles
}