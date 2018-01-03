package expressions

import "github.com/cdstelly/nugget/NTypes"

//todo: big todo: move from interface on GetResults to a custom nugget baseType -- we only want to deal with types we implcitly know how to handle
type BaseAction interface {
	BeenExecuted() bool
	DependencySatisfied() bool
	SetDependency(action BaseAction)
	Execute()
	GetResults() interface{}
	SetFilters([]NTypes.Filter)
}