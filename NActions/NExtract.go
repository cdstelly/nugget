package NActions

//don't export this type so that we can force users of it to use the 'New' method, thereby initializing values
type ExtractNTFS struct {
	executed  bool
	dependsOn BaseAction

	NTFSImageLocation string
}

func NewNExtractNTFS() ExtractNTFS {
	n := ExtractNTFS{}
	n.dependsOn = nil
	n.executed = false
	return n
}

func (na *ExtractNTFS) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractNTFS) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractNTFS) SetDependency(action BaseAction) {
	na.dependsOn = nil
}

func (na *ExtractNTFS) Exeucte() {
	na.executed = true

}