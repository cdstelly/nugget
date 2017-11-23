package expressions

import "github.com/cdstelly/nugget/NTypes"

type BaseAction interface {
	BeenExecuted() bool
	DependencySatisfied() bool
	SetDependency(action BaseAction)
	Execute()
	GetResults() interface{}
	SetFilters([]NTypes.Filter)
}