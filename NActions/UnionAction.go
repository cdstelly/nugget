package NActions

import (
	"../NTypes"
)

type UnionAction struct {
	executed  bool
	dependsOn BaseAction
	filters   []NTypes.Filter

	results      []string
	VariableList []string //the stuff we get in to union against
}

func (na *UnionAction) BeenExecuted() bool {
	return na.executed
}

func (na *UnionAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *UnionAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}
func (na *UnionAction) Execute() {
	if na.dependsOn != nil {
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	resultSet := []string{}

	operateOn := na.dependsOn.GetResults()

	//todo what about sha1 and sha256..
	//todo optimize this n^2 comparison loop..
	if strList, ok := operateOn.([]NTypes.MD5); ok {
//		fmt.Println("going to execute union against a hash..")

		for _, stringFromVar := range na.VariableList {
			for _, stringFromDep := range strList {
				//fmt.Println("Comparing: ", stringFromVar, stringFromDep)
				if stringFromVar == stringFromDep.Digest {
					if fileData, weHaveFileInfo := stringFromDep.HashOf.(NTypes.FileInfo); weHaveFileInfo {
						fn := fileData.Filenames[0]
						resultSet = append(resultSet, fn + "\t" + stringFromDep.Digest)
					} else {
						resultSet = append(resultSet, stringFromDep.Digest)
					}
				}
			}
		}

		//todo add unions for all types, but do it at the type level 
	}

	na.results = resultSet
	na.executed = true
}

func (na *UnionAction) GetResults() interface{} {
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *UnionAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}
