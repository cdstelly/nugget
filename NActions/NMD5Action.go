package NActions

import "fmt"
import (
	"../NTypes"
	"crypto/md5"
)

type MD5Action struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	results []NTypes.MD5
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
func (na *MD5Action) Execute() {
	if na.dependsOn != nil {
		//fmt.Println("md5 has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {

			na.dependsOn.Execute()
		}
	}
	//fmt.Println("going to execute md5..")

	operateOn := na.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes.FileInfo); ok {
		var files []NTypes.FileInfo
		files = operateOn.([]NTypes.FileInfo)
		for _,file := range files {
			hasher := md5.New()
			hasher.Write(file.GetFileData())
			myhash := fmt.Sprintf("%x", hasher.Sum(nil))
			//fmt.Println("fn:" + file.Filenames[0] + "\tmd5: " + myhash)
			na.results = append(na.results, NTypes.MD5{Digest:myhash})
		}
	}
	na.executed = true
}

func (na *MD5Action) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *MD5Action) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}