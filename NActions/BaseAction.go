package NActions

type BaseAction interface {
	BeenExecuted() bool
	DependencySatisfied() bool
	SetDependency(action BaseAction)
	Execute()
	GetResults() interface{}
}