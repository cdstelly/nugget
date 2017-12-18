package extractors

import (
	"github.com/cdstelly/nugget/NTypes"

	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"fmt"
	"strings"
	"github.com/cdstelly/nugget/expressions/transforms"
)

type ExtractPCAP struct {
	executed  bool
	dependsOn expressions.BaseAction
	filters []NTypes.Filter

	PCAPLocation string
	handle *pcap.Handle
	Packets []NTypes.NPacket
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

func (na *ExtractPCAP) GetPackets() []NTypes.NPacket{
	//fmt.Println("getting packets")
	var err error

	//get bpfs and apply them
	var myBPF string
	for _, f := range na.filters {
		//fmt.Println("Found a filter: ", f)
		switch f.Field {
		case "packetfilter":
			if f.Op == "==" {
				myBPF = f.Value
				myBPF = strings.Trim(myBPF, "\"")
				//fmt.Println("Set the bnf filter value to: " + myBPF)
			} else {
				fmt.Println("Error: BNF only supports equality at the moment.")
			}
		default:
			fmt.Println("Error: PCAP Parser was unable to understand filter: ", f.Field)
		}
	}
	//fmt.Println("pcap location: " + na.PCAPLocation)
	na.handle, err = pcap.OpenOffline(na.PCAPLocation)
	if err != nil {
		fmt.Println("Error reading pcap file: " , err)
		//panic(err)
	}
	err = na.handle.SetBPFFilter(myBPF)
	if err != nil {
		fmt.Println("Error setting BPF: " , err)
		panic(err)
	}

	packetSource := gopacket.NewPacketSource(na.handle, na.handle.LinkType())

	for p := range packetSource.Packets() {
		na.Packets = append(na.Packets, NTypes.NPacket{p})
	}
	na.executed = true
	return na.Packets
}

func (na *ExtractPCAP) GetResults() interface{}{
	if na.executed == false {
		na.Execute()
	}
	//fmt.Println(len(na.Packets) ," ", reflect.TypeOf(na.Packets))
	return na.Packets
}

func (na *ExtractPCAP) SetFilters(filters []NTypes.Filter) {
	na.filters = filters
}
