package NActions

import "fmt"
import (
	"../NTypes"
	"crypto/sha256"
)

type SHA256Action struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	results []NTypes.SHA256
}

func (na *SHA256Action) BeenExecuted() bool {
	return na.executed
}

func (na *SHA256Action) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *SHA256Action) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (na *SHA256Action) Execute() {
	if na.dependsOn != nil {
		fmt.Println("SHA256Action has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	fmt.Println("going to execute sha256..")

	operateOn := na.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes.FileInfo); ok {
		//we have a file to operate on
		var files []NTypes.FileInfo
		files = operateOn.([]NTypes.FileInfo)
		for index,file := range files {
			hasher := sha256.New()
			fn := GetAFilename(file)
			hasher.Write([]byte(fn))
			myhash := fmt.Sprintf("%x", hasher.Sum(nil))
			println("index: ", index, " file: ", file.Filenames, " md5: ", myhash)
			na.results = append(na.results, NTypes.SHA256{Digest:myhash})
		}
	}
	na.executed = true
}

func (na *SHA256Action) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *SHA256Action) SetFilters(filters []NTypes.Filter) {
	na.filters = filters
}