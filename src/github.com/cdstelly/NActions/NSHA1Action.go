package NActions

import (
	"github.com/cdstelly/NTypes"
	"fmt"
	"crypto/sha1"
)

type SHA1Action struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter


	target []NTypes.FileInfo
	results []NTypes.SHA1
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

func (na *SHA1Action) Execute() {
	if na.dependsOn != nil {
		//fmt.Println("sha1 has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	//fmt.Println("going to execute sha1..")

	if na.dependsOn == nil {
		fmt.Println("Error! Was not able to compute SHA1 because dependency is nil")
		return
	}

	operateOn := na.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes.FileInfo); ok {
		//we have a file to operate on
		var files []NTypes.FileInfo
		files = operateOn.([]NTypes.FileInfo)
		for _,file := range files {
			hasher := sha1.New()
			fn := GetAFilename(file)
			hasher.Write([]byte(fn))
			myhash := fmt.Sprintf("%x", hasher.Sum(nil))
			//println("index: ", index, " file: ", file.Filenames, " sha1: ", myhash)
			na.results = append(na.results, NTypes.SHA1{myhash})
		}
	}
	na.executed = true
}

func (na *SHA1Action) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *SHA1Action) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}