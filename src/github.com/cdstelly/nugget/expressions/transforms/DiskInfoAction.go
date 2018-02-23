package expressions


import (
	"github.com/cdstelly/nugget/NTypes"
	"net/rpc"
	"log"
	"strings"
//	"fmt"
)

type DiskInfoAction struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	results NTypes.DiskImageInfo

	beenUploaded bool
}

func (na *DiskInfoAction) BeenExecuted() bool {
	return na.executed
}

func (na *DiskInfoAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *DiskInfoAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}
func (na *DiskInfoAction) Execute() {
	if na.dependsOn != nil {
		//fmt.Println("disk image info has a dependency which hasn't been met..")
		if na.dependsOn.BeenExecuted() == false {
			na.dependsOn.Execute()
		}
	}
	//fmt.Println("going to execute disk image info..")
	//operateOn := na.dependsOn.GetResults()  //should be a raw disk image

	//todo: uploadimage is not necessary since extractfs should occur before this
	na.uploadImageToTSK()
	diskImageInfo := getImageInfoFromTSK()
	//fmt.Println(diskImageInfo)

	theInfo := strings.Split(diskImageInfo,"\n")
	var theInfoMap map[string]string
	theInfoMap = make(map[string]string)

	//for every line of text
	for _,ti := range theInfo {
		kv := strings.Split(ti, ":")
		if len(kv)>1 {
			key := kv[0]
			val := kv[1]
			theInfoMap[key] = val
		}
	}
	//todo how to handle new versions of FSSTAT?
	myDiskImageInfo := NTypes.DiskImageInfo{}
	myDiskImageInfo.FSType = theInfoMap["File System Type"]
	myDiskImageInfo.ClusterRange = theInfoMap["Total Cluster Range"]
	myDiskImageInfo.ClusterSize = theInfoMap["Cluster Size"]
	myDiskImageInfo.FirstCluster = theInfoMap["First Cluster of MFT"]
	myDiskImageInfo.OEMName = theInfoMap["OEM Name"]
	myDiskImageInfo.SerialNumber = theInfoMap["Volume Serial Number"]
	myDiskImageInfo.Version = theInfoMap["Version"]

	na.results = myDiskImageInfo
	na.executed = true
}

func (na *DiskInfoAction) GetResults() interface{}{
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.results
}

func (na *DiskInfoAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}


func getImageInfoFromTSK() string {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	//load some data into tsk memory
	args := &NTypes.NugArg{[]byte(""),""}
	var reply string
	err = client.Call("NugTSK.ExecImageInfo", args, &reply)
	if err != nil {
		log.Fatal("tsk load error:", err)
	}
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	return reply
}

func (na *DiskInfoAction) uploadImageToTSK() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//load some data into tsk memory
	args := &NTypes.NugArg{[]byte("test"),""}
	var reply string
	err = client.Call("NugTSK.LoadData", args, &reply)
	if err != nil {
		log.Fatal("tsk load error:", err)
	}
	//fmt.Printf("tsk: %s=%s\n", string(args.TheData), reply)
	na.beenUploaded = true
}