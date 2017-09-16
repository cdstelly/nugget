package NActions

import (
	"../NTypes"

	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"fmt"
	"strings"
)

//don't export this type so that we can force users of it to use the 'New' method, thereby initializing values
type ExtractPCAP struct {
	executed  bool
	dependsOn BaseAction
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

func (na *ExtractPCAP) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (na *ExtractPCAP) Execute() {
	na.Packets = na.GetPackets()
	na.executed = true
}

func (na *ExtractPCAP) GetPackets() []NTypes.NPacket{
	fmt.Println("getting packets")
	var err error

	//get bpfs and apply them
	var myBPF string
	for _, f := range na.filters {
		fmt.Println("Found a filter: ", f)
		switch f.Field {
		case "packetfilter":
			if f.Op == "==" {
				myBPF = f.Value
				myBPF = strings.Trim(myBPF, "\"")
				fmt.Println("Set the bnf filter value to: ", myBPF )
			} else {
				fmt.Println("Error: BNF only supports equality at the moment.")
			}
		default:
			fmt.Println("Error: PCAP Parser was unable to understand filter: ", f.Field)
		}
	}

	na.handle, err = pcap.OpenOffline(na.PCAPLocation)
	if err != nil {
		panic(err)
	}
	err = na.handle.SetBPFFilter(myBPF)
	if err != nil {
		panic(err)
	}

	packetSource := gopacket.NewPacketSource(na.handle, na.handle.LinkType())

	// asynchronously read the packet source
	for p := range packetSource.Packets() {
		na.Packets = append(na.Packets, NTypes.NPacket{p})
	}
	na.executed = true
	fmt.Println(len(na.Packets))
	return na.Packets
}

func (na *ExtractPCAP) GetResults() interface{}{
	if na.executed == false {
		na.Execute()
	}
	return na.Packets
}

func (na *ExtractPCAP) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	//na.executed = false
	na.filters = filters
}

// want to make filter an action

// so that when we do 'x = var | filter x = "y"
// the action 'filter..' will pull in the results from executing 'var', then try and apply the filter, then return the results and store into 'x'
