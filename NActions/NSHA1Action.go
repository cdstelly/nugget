package NActions

import (
	"../NTypes"
	"fmt"
	"crypto/sha1"
)

type SHA1Action struct {
	executed  bool
	dependsOn BaseAction


	target []NTypes.FileInfo
	results []NTypes.SHA1

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

func (na *SHA1Action) Execute() {
	fmt.Println("going to execute sha1 on results from dependency: ", na.dependsOn)
	if na.dependsOn.BeenExecuted() == false {
		na.dependsOn.Execute()
	}

	operateOn := na.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes.FileInfo); ok {
		//we have a file to operate on
		var files []NTypes.FileInfo
		files = operateOn.([]NTypes.FileInfo)
		for index,file := range files {
			hasher := sha1.New()
			fn := GetAFilename(file)
			hasher.Write([]byte(fn))
			myhash := fmt.Sprintf("%x\n", hasher.Sum(nil))
			println("index: ", index, " file: ", file.Filenames, " sha1: ", myhash)
			na.results = append(na.results, NTypes.SHA1{myhash})
		}
	}
}


func (na *SHA1Action) GetResults() interface{}{
	return na.results
}