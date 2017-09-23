package NActions

import "fmt"
import (
	"../NTypes"
	"sort"
)

type SortAction struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	SortField string

	results interface{}
}

func (na *SortAction) BeenExecuted() bool {
	return na.executed
}

func (na *SortAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *SortAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}
func (na *SortAction) Execute() {
	if na.dependsOn != nil {
		fmt.Println("sortaction has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	fmt.Println("going to execute sort..")

	operateOn := na.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes.FileInfo); ok {
		var files []NTypes.FileInfo
		files = operateOn.([]NTypes.FileInfo)
		if na.SortField == "ctime" {
			sort.Sort(NTypes.ByCtime(files))
		}
		if na.SortField == "mtime" {
			sort.Sort(NTypes.ByMtime(files))
		}
		if na.SortField == "etime" {
			sort.Sort(NTypes.ByEtime(files))
		}
		if na.SortField == "atime" {
			sort.Sort(NTypes.ByAtime(files))
		}
		if na.SortField == "filename" {
			sort.Sort(NTypes.ByFilename(files))
		}
		if na.SortField == "filesize" {
			sort.Sort(NTypes.ByFilesize(files))
		}
		na.results = files
	}
	na.executed = true
}

func (na *SortAction) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *SortAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}