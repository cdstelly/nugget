package NActions

type SHA1Action struct {
	executed  bool
	dependsOn BaseAction
}

func NewSHA1Action(dep BaseAction) SHA1Action {
	n := SHA1Action{}
	n.dependsOn = dep
	n.executed = false
	return n
}

func (na *SHA1Action) BeenExecuted() bool {
	return na.executed
}

func (na *SHA1Action) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *SHA1Action) SetDependency(action BaseAction) {
	na.dependsOn = action
}