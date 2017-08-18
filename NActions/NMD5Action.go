package NActions

type MD5Action struct {
	executed  bool
	dependsOn BaseAction
}

func NewMD5Action(dep BaseAction) MD5Action {
	n := MD5Action{}
	n.dependsOn = dep
	n.executed = false
	return n
}

func (na *MD5Action) BeenExecuted() bool {
	return na.executed
}

func (na *MD5Action) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *MD5Action) SetDependency(action BaseAction) {
	na.dependsOn = action
}
