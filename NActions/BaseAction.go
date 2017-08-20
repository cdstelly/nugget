package NActions

import "../NTypes"

type BaseAction interface {
	BeenExecuted() bool
	DependencySatisfied() bool
	SetDependency(action BaseAction)
	Execute()
	GetResults() interface{}
	SetFilters([]NTypes.Filter)
}