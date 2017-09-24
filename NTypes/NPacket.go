package NTypes

import (
	"github.com/google/gopacket"
	"fmt"
)

type NPacket struct {
	Pkt gopacket.Packet
}

func (p NPacket) String() string {

	applicationLayer := p.Pkt.ApplicationLayer()
	if applicationLayer != nil {
		return string(applicationLayer.Payload())
	}
	return fmt.Sprintf( "unable to decode the packet" )
}