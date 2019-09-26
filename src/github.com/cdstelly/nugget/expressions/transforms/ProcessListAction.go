package expressions

import (
	"github.com/cdstelly/nugget/NTypes"
	"log"
	"net/rpc"
)

type ProcessListAction struct {
	executed  bool
	dependsOn BaseAction
	filters   []NTypes.Filter

	results NTypes.ProcessList

	beenUploaded bool
}

func (na *ProcessListAction) BeenExecuted() bool {
	return na.executed
}

func (na *ProcessListAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *ProcessListAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (na *ProcessListAction) Execute() {
	if na.dependsOn != nil {
		//fmt.Println("process list has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	//fmt.Println("going to execute process list..")
	//operateOn := na.dependsOn.GetResults()  //should be a raw disk image

	//na.uploadImageToVOL()
	procList := getPSListFromVOL()
	na.results = NTypes.ProcessList{Processes: procList}
	na.executed = true
}

func (na *ProcessListAction) GetResults() interface{} {
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *ProcessListAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}

func getPSListFromVOL() string {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//load some data into tsk memory
	args := &NTypes.NugArg{[]byte(""), ""}
	var reply string
	err = client.Call("NugVol.PSList", args, &reply)
	if err != nil {
		log.Fatal("[!] Volatility Error:", reply, err)
	}

	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	return reply
}
