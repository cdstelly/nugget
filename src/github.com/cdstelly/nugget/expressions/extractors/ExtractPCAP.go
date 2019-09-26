package extractors

import (
	"encoding/gob"
	"github.com/cdstelly/nugget/NTypes"
	"log"
	"net/rpc"

	"github.com/cdstelly/nugget/expressions/transforms"
	"github.com/google/gopacket/pcap"
)

type ExtractPCAP struct {
	executed  bool
	dependsOn expressions.BaseAction
	filters   []NTypes.Filter

	PcapClient   *rpc.Client
	beenUploaded bool

	Location string
	handle   *pcap.Handle
	Packets  []NTypes.NPacket
}

func (na *ExtractPCAP) BeenExecuted() bool {
	return na.executed
}

func (na *ExtractPCAP) DependencySatisfied() bool {
	return true //extractions don't depend on any other actions to execute
}

func (na *ExtractPCAP) SetDependency(action expressions.BaseAction) {
	na.dependsOn = action
}

func (na *ExtractPCAP) Execute() {
	na.Packets = na.GetPackets()
	na.executed = true
}

func (na *ExtractPCAP) ExtractPacketsFromPcap() []NTypes.NPacket {

	gob.Register(NTypes.NPacket{})

	if na.beenUploaded == false {
		na.UploadData()
	}

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2001")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &NTypes.NuggetPCAPArgs{nil, "", nil}
	//var reply []NTypes.NPacket
	var reply []string
	err = client.Call("NugPcap.GetPackets", args, &reply)
	if err != nil {
		log.Fatal("[!] Pcap RPC error:", err)
	}
	log.Println("temp response: ", reply)
	return nil

}

func (na *ExtractPCAP) GetPackets() []NTypes.NPacket {
	log.Println("[-] Obtaining Packets..")

	// upload data
	if na.beenUploaded == false {
		na.UploadData()
	}

	// send filters to the RPC processor
	na.UploadFilters()

	// get packets
	packetizedResponse := na.ExtractPacketsFromPcap()
	log.Println(packetizedResponse)

	na.executed = true
	return na.Packets
}

func (na *ExtractPCAP) GetResults() interface{} {
	if na.executed == false {
		na.Execute()
	}
	//fmt.Println(len(na.Packets) ," ", reflect.TypeOf(na.Packets))
	return na.Packets
}

func (na *ExtractPCAP) SetFilters(filters []NTypes.Filter) {
	na.filters = filters
}

func (na *ExtractPCAP) UploadData() {
	log.Println("[-] Uploading pcap information to RPC", na.Location)
	//load some data into RPC PCAP processor memory
	args := &NTypes.NuggetPCAPArgs{PcapLocation: na.Location}
	var reply string
	err := na.PcapClient.Call("NugPcap.LoadData", args, &reply)
	if err != nil {
		log.Fatal("PCAP RPC (local) error - priming data store:", err)
	}
	log.Printf("pcap response: %s=%s\n", string(args.TheData), reply)
	na.beenUploaded = true
}

func (na *ExtractPCAP) UploadFilters() {
	//load give filters into RPC PCAP processor memory
	args := &NTypes.NuggetPCAPArgs{Filters: na.filters}
	var reply string
	err := na.PcapClient.Call("NugPcap.LoadFilters", args, &reply)
	if err != nil {
		log.Fatal("PCAP RPC (local) error - uploading filter information to pcap processor:", err)
	}

	na.beenUploaded = true
}
