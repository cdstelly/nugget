package extractors

import (
	"github.com/cdstelly/nugget/NTypes"
	"github.com/cdstelly/nugget/expressions/transforms"
	"log"
	"net/rpc"
)

type ExtractMemory struct {
	executed  bool
	dependsOn expressions.BaseAction
	filters   []NTypes.Filter

	Location     string
	beenUploaded bool

	VolClient *rpc.Client
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

func (na *ExtractMemory) GetResults() interface{} {
	if na.executed == false {
		na.Execute()
	}
	return na.Location
}

func (na *ExtractMemory) SetFilters(filters []NTypes.Filter) {
	na.filters = filters
}

func (na *ExtractMemory) UploadData() {

	//load some data into tsk memory
	args := &NTypes.NugArg{[]byte(na.Location), ""}
	var reply string
	volUploadError := na.VolClient.Call("NugVol.LoadData", args, &reply)
	if volUploadError != nil {
		log.Fatal("volatility load error:", volUploadError)
	}
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	na.beenUploaded = true
}
