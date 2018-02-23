package extractors

import (
	"github.com/cdstelly/nugget/NTypes"
	"github.com/cdstelly/nugget/expressions/transforms"
	"net/rpc"
	"log"
)

type ExtractMemory struct {
	executed  bool
	dependsOn expressions.BaseAction
	filters []NTypes.Filter

	Location string
	beenUploaded bool
}

func (na *ExtractMemory) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractMemory) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractMemory) SetDependency(action expressions.BaseAction) {
	na.dependsOn = action
}

func (na *ExtractMemory) Execute() {
	if na.beenUploaded == false {
		na.UploadData()
	}
	na.executed = true
}

func (na *ExtractMemory) GetResults() interface{}{
	if na.executed == false {
		na.Execute()
	}
	return na.Location
}

func (na *ExtractMemory) SetFilters(filters []NTypes.Filter) {
	na.filters = filters
}

func (na *ExtractMemory) UploadData() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2002")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//load some data into tsk memory
	args := &NTypes.NugArg{[]byte(na.Location),""}
	var reply string
	err = client.Call("NugVol.LoadData", args, &reply)
	if err != nil {
		log.Fatal("tsk load error:", err)
	}
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	na.beenUploaded = true
}