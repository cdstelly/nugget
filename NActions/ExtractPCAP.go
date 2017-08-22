package NActions

import (
	"../NTypes"

	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
)

//don't export this type so that we can force users of it to use the 'New' method, thereby initializing values
type ExtractPCAP struct {
	executed  bool
	dependsOn BaseAction
	filters []NTypes.Filter

	PCAPLocation string
	handle *pcap.Handle
	Packets []gopacket.Packet
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

func (na *ExtractPCAP) GetPackets() []gopacket.Packet{
	var err error

	na.handle, err = pcap.OpenOffline(na.PCAPLocation)
	checkError(err)

	packetSource := gopacket.NewPacketSource(na.handle, na.handle.LinkType())

	// asynchronously read the packet source
	for p := range packetSource.Packets() {
		na.Packets = append(na.Packets, p)
	}
	na.executed = true
	return na.Packets
}

func (na *ExtractPCAP) GetResults() interface{}{
	return na.Packets
}

func (na *ExtractPCAP) SetFilters(filters []NTypes.Filter) {
	na.filters = filters
}